package v20221222

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type QueryInstanceConsumeRequest struct {
	*ksyunhttp.BaseRequest
	StartDay    *string `json:"StartDay,omitempty" name:"StartDay"`
	EndDay      *string `json:"EndDay,omitempty" name:"EndDay"`
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	Page        *int    `json:"Page,omitempty" name:"Page"`
	Size        *int    `json:"Size,omitempty" name:"Size"`
}

func (r *QueryInstanceConsumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryInstanceConsumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryInstanceConsumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryInstanceConsumeResponse struct {
	*ksyunhttp.BaseResponse
	Status    *int    `json:"status" name:"status"`
	RequestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		Page  *int `json:"Page" name:"Page"`
		Size  *int `json:"Size" name:"Size"`
		Total *int `json:"Total" name:"Total"`
		Bills []struct {
			AliasName               *string   `json:"aliasName" name:"aliasName"`
			Email                   *string   `json:"email" name:"email"`
			MembershipGroup         *string   `json:"membershipGroup" name:"membershipGroup"`
			UserId                  *int      `json:"userId" name:"userId"`
			UserName                *string   `json:"userName" name:"userName"`
			SellerCompanyName       *string   `json:"sellerCompanyName" name:"sellerCompanyName"`
			BillMonth               *int      `json:"billMonth" name:"billMonth"`
			CustomerBillMonth       *int      `json:"customerBillMonth" name:"customerBillMonth"`
			CurrencyCode            *string   `json:"currencyCode" name:"currencyCode"`
			CurrencyInfo            *string   `json:"currencyInfo" name:"currencyInfo"`
			ExchangeRate            *string   `json:"exchangeRate" name:"exchangeRate"`
			BillDay                 *int      `json:"billDay" name:"billDay"`
			Id                      *string   `json:"id" name:"id"`
			FinanceUnitName         *string   `json:"financeUnitName" name:"financeUnitName"`
			BillStartTime           *string   `json:"billStartTime" name:"billStartTime"`
			BillEndTime             *string   `json:"billEndTime" name:"billEndTime"`
			InstanceId              *string   `json:"instanceId" name:"instanceId"`
			InstanceName            *string   `json:"instanceName" name:"instanceName"`
			ProductTypeId           *int      `json:"productTypeId" name:"productTypeId"`
			ProductTypeName         *string   `json:"productTypeName" name:"productTypeName"`
			ProductGroupId          *int      `json:"productGroupId" name:"productGroupId"`
			ProductGroupName        *string   `json:"productGroupName" name:"productGroupName"`
			PayType                 *int      `json:"payType" name:"payType"`
			BillRealAmount          *string   `json:"billRealAmount" name:"billRealAmount"`
			OriginalAmount          *string   `json:"originalAmount" name:"originalAmount"`
			RegionCode              *string   `json:"regionCode" name:"regionCode"`
			ProjectId               *int      `json:"projectId" name:"projectId"`
			ProjectName             *string   `json:"projectName" name:"projectName"`
			RegionName              *string   `json:"regionName" name:"regionName"`
			BillType                *int      `json:"billType" name:"billType"`
			BillTypeName            *string   `json:"billTypeName" name:"billTypeName"`
			PayTypeName             *string   `json:"payTypeName" name:"payTypeName"`
			BillDetailType          *int      `json:"billDetailType" name:"billDetailType"`
			BillDetailTypeName      *string   `json:"billDetailTypeName" name:"billDetailTypeName"`
			Duration                *string   `json:"duration" name:"duration"`
			DurationNumber          *int      `json:"durationNumber" name:"durationNumber"`
			RuleRemark              *string   `json:"ruleRemark" name:"ruleRemark"`
			AvailabilityZone        *string   `json:"availabilityZone" name:"availabilityZone"`
			Discount                *string   `json:"discount" name:"discount"`
			Cash                    *string   `json:"cash" name:"cash"`
			Reward                  *string   `json:"reward" name:"reward"`
			Vouchers                *string   `json:"vouchers" name:"vouchers"`
			CloudTicketDenomination *string   `json:"cloudTicketDenomination" name:"cloudTicketDenomination"`
			CloudTicketCost         *string   `json:"cloudTicketCost" name:"cloudTicketCost"`
			ServiceBegionTime       *string   `json:"serviceBegionTime" name:"serviceBegionTime"`
			ConfigInfo              []*string `json:"ConfigInfo" name:"ConfigInfo"`
			PriceFactorInfo         []*string `json:"PriceFactorInfo" name:"PriceFactorInfo"`
			ExtraInfo               []*string `json:"ExtraInfo" name:"ExtraInfo"`
			ResourceDeductionInfo   []*string `json:"ResourceDeductionInfo" name:"ResourceDeductionInfo"`
			TagInfo                 []*string `json:"TagInfo" name:"TagInfo"`
		} `json:"Bills" name:"Bills"`
	} `json:"Data"`
	Error *string `json:"error" name:"error"`
	Ok    *bool   `json:"ok" name:"ok"`
}

func (r *QueryInstanceConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryInstanceConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProjectConsumeRequest struct {
	*ksyunhttp.BaseRequest
	StartDay *string `json:"StartDay,omitempty" name:"StartDay"`
	EndDay   *string `json:"EndDay,omitempty" name:"EndDay"`
	Page     *int    `json:"Page,omitempty" name:"Page"`
	Size     *int    `json:"Size,omitempty" name:"Size"`
}

func (r *QueryProjectConsumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProjectConsumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryProjectConsumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryProjectConsumeResponse struct {
	*ksyunhttp.BaseResponse
	Status    *int    `json:"status" name:"status"`
	RequestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		Page  *int `json:"Page" name:"Page"`
		Size  *int `json:"Size" name:"Size"`
		Total *int `json:"Total" name:"Total"`
		Bills []struct {
			AliasName               *string `json:"aliasName" name:"aliasName"`
			Email                   *string `json:"email" name:"email"`
			MembershipGroup         *string `json:"membershipGroup" name:"membershipGroup"`
			UserId                  *int    `json:"userId" name:"userId"`
			UserName                *string `json:"userName" name:"userName"`
			SellerCompanyName       *string `json:"sellerCompanyName" name:"sellerCompanyName"`
			BillMonth               *string `json:"billMonth" name:"billMonth"`
			CustomerBillMonth       *int    `json:"customerBillMonth" name:"customerBillMonth"`
			CurrencyCode            *string `json:"currencyCode" name:"currencyCode"`
			CurrencyInfo            *string `json:"currencyInfo" name:"currencyInfo"`
			ExchangeRate            *string `json:"exchangeRate" name:"exchangeRate"`
			BillDay                 *int    `json:"billDay" name:"billDay"`
			ProjectId               *int    `json:"projectId" name:"projectId"`
			ProjectName             *string `json:"projectName" name:"projectName"`
			PrePayCost              *string `json:"prePayCost" name:"prePayCost"`
			PrePayRefundAmount      *string `json:"prePayRefundAmount" name:"prePayRefundAmount"`
			PrePayBillsAmount       *string `json:"prePayBillsAmount" name:"prePayBillsAmount"`
			RealtimeCost            *string `json:"realtimeCost" name:"realtimeCost"`
			RealtimeAdjustAmount    *string `json:"realtimeAdjustAmount" name:"realtimeAdjustAmount"`
			RealtimeBillsAmount     *string `json:"realtimeBillsAmount" name:"realtimeBillsAmount"`
			PostPayCost             *string `json:"postPayCost" name:"postPayCost"`
			PostPayAdjustAmount     *string `json:"postPayAdjustAmount" name:"postPayAdjustAmount"`
			PostPayBillsAmount      *string `json:"postPayBillsAmount" name:"postPayBillsAmount"`
			BillFinalAmount         *string `json:"billFinalAmount" name:"billFinalAmount"`
			Cash                    *string `json:"cash" name:"cash"`
			Reward                  *string `json:"reward" name:"reward"`
			Vouchers                *string `json:"vouchers" name:"vouchers"`
			CloudTicketDenomination *string `json:"cloudTicketDenomination" name:"cloudTicketDenomination"`
		} `json:"Bills" name:"Bills"`
	} `json:"Data"`
	Error *string `json:"error" name:"error"`
	Ok    *bool   `json:"ok" name:"ok"`
}

func (r *QueryProjectConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProjectConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductConsumeRequest struct {
	*ksyunhttp.BaseRequest
	StartDay *string `json:"StartDay,omitempty" name:"StartDay"`
	EndDay   *string `json:"EndDay,omitempty" name:"EndDay"`
	Page     *int    `json:"Page,omitempty" name:"Page"`
	Size     *int    `json:"Size,omitempty" name:"Size"`
}

func (r *QueryProductConsumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductConsumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryProductConsumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryProductConsumeResponse struct {
	*ksyunhttp.BaseResponse
	Status    *int    `json:"status" name:"status"`
	RequestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		Page  *int `json:"Page" name:"Page"`
		Size  *int `json:"Size" name:"Size"`
		Total *int `json:"Total" name:"Total"`
		Bills []struct {
			AliasName               *string `json:"aliasName" name:"aliasName"`
			Email                   *string `json:"email" name:"email"`
			MembershipGroup         *string `json:"membershipGroup" name:"membershipGroup"`
			UserId                  *int    `json:"userId" name:"userId"`
			UserName                *string `json:"userName" name:"userName"`
			SellerCompanyName       *string `json:"sellerCompanyName" name:"sellerCompanyName"`
			BillMonth               *string `json:"billMonth" name:"billMonth"`
			CustomerBillMonth       *int    `json:"customerBillMonth" name:"customerBillMonth"`
			CurrencyCode            *string `json:"currencyCode" name:"currencyCode"`
			CurrencyInfo            *string `json:"currencyInfo" name:"currencyInfo"`
			ExchangeRate            *string `json:"exchangeRate" name:"exchangeRate"`
			BillDay                 *int    `json:"billDay" name:"billDay"`
			ProductGroupId          *int    `json:"productGroupId" name:"productGroupId"`
			ProductGroupName        *string `json:"productGroupName" name:"productGroupName"`
			PrePayCost              *string `json:"prePayCost" name:"prePayCost"`
			PrePayRefundAmount      *string `json:"prePayRefundAmount" name:"prePayRefundAmount"`
			PrePayBillsAmount       *string `json:"prePayBillsAmount" name:"prePayBillsAmount"`
			RealtimeCost            *string `json:"realtimeCost" name:"realtimeCost"`
			RealtimeAdjustAmount    *string `json:"realtimeAdjustAmount" name:"realtimeAdjustAmount"`
			RealtimeBillsAmount     *string `json:"realtimeBillsAmount" name:"realtimeBillsAmount"`
			PostPayCost             *string `json:"postPayCost" name:"postPayCost"`
			PostPayAdjustAmount     *string `json:"postPayAdjustAmount" name:"postPayAdjustAmount"`
			PostPayBillsAmount      *string `json:"postPayBillsAmount" name:"postPayBillsAmount"`
			BillFinalAmount         *string `json:"billFinalAmount" name:"billFinalAmount"`
			Cash                    *string `json:"cash" name:"cash"`
			Reward                  *string `json:"reward" name:"reward"`
			Vouchers                *string `json:"vouchers" name:"vouchers"`
			CloudTicketDenomination *string `json:"cloudTicketDenomination" name:"cloudTicketDenomination"`
		} `json:"Bills" name:"Bills"`
	} `json:"Data"`
	Error *string `json:"error" name:"error"`
	Ok    *bool   `json:"ok" name:"ok"`
}

func (r *QueryProductConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryFinanceUnitConsumeRequest struct {
	*ksyunhttp.BaseRequest
	StartDay *string `json:"StartDay,omitempty" name:"StartDay"`
	EndDay   *string `json:"EndDay,omitempty" name:"EndDay"`
	Page     *int    `json:"Page,omitempty" name:"Page"`
	Size     *int    `json:"Size,omitempty" name:"Size"`
}

func (r *QueryFinanceUnitConsumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryFinanceUnitConsumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryFinanceUnitConsumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryFinanceUnitConsumeResponse struct {
	*ksyunhttp.BaseResponse
	Status    *int    `json:"status" name:"status"`
	RequestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		Page  *int `json:"Page" name:"Page"`
		Size  *int `json:"Size" name:"Size"`
		Total *int `json:"Total" name:"Total"`
		Bills []struct {
			AliasName               *string `json:"aliasName" name:"aliasName"`
			Email                   *string `json:"email" name:"email"`
			MembershipGroup         *string `json:"membershipGroup" name:"membershipGroup"`
			UserId                  *int    `json:"userId" name:"userId"`
			UserName                *string `json:"userName" name:"userName"`
			SellerCompanyName       *string `json:"sellerCompanyName" name:"sellerCompanyName"`
			BillMonth               *string `json:"billMonth" name:"billMonth"`
			CustomerBillMonth       *int    `json:"customerBillMonth" name:"customerBillMonth"`
			CurrencyCode            *string `json:"currencyCode" name:"currencyCode"`
			CurrencyInfo            *string `json:"currencyInfo" name:"currencyInfo"`
			ExchangeRate            *string `json:"exchangeRate" name:"exchangeRate"`
			BillDay                 *int    `json:"billDay" name:"billDay"`
			FinanceUnitName         *string `json:"financeUnitName" name:"financeUnitName"`
			PrePayCost              *string `json:"prePayCost" name:"prePayCost"`
			PrePayRefundAmount      *string `json:"prePayRefundAmount" name:"prePayRefundAmount"`
			PrePayBillsAmount       *string `json:"prePayBillsAmount" name:"prePayBillsAmount"`
			RealtimeCost            *string `json:"realtimeCost" name:"realtimeCost"`
			RealtimeAdjustAmount    *string `json:"realtimeAdjustAmount" name:"realtimeAdjustAmount"`
			RealtimeBillsAmount     *string `json:"realtimeBillsAmount" name:"realtimeBillsAmount"`
			PostPayCost             *string `json:"postPayCost" name:"postPayCost"`
			PostPayAdjustAmount     *string `json:"postPayAdjustAmount" name:"postPayAdjustAmount"`
			PostPayBillsAmount      *string `json:"postPayBillsAmount" name:"postPayBillsAmount"`
			BillFinalAmount         *string `json:"billFinalAmount" name:"billFinalAmount"`
			Cash                    *string `json:"cash" name:"cash"`
			Reward                  *string `json:"reward" name:"reward"`
			Vouchers                *string `json:"vouchers" name:"vouchers"`
			CloudTicketDenomination *string `json:"cloudTicketDenomination" name:"cloudTicketDenomination"`
		} `json:"Bills" name:"Bills"`
	} `json:"Data"`
	Error *string `json:"error" name:"error"`
	Ok    *bool   `json:"ok" name:"ok"`
}

func (r *QueryFinanceUnitConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryFinanceUnitConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryFinanceUnitConsumeOfMonthRequest struct {
	*ksyunhttp.BaseRequest
	CustomerBillMonth *string `json:"CustomerBillMonth,omitempty" name:"CustomerBillMonth"`
	Page              *int    `json:"Page,omitempty" name:"Page"`
	Size              *int    `json:"Size,omitempty" name:"Size"`
}

func (r *QueryFinanceUnitConsumeOfMonthRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryFinanceUnitConsumeOfMonthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryFinanceUnitConsumeOfMonthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryFinanceUnitConsumeOfMonthResponse struct {
	*ksyunhttp.BaseResponse
	Status    *int    `json:"status" name:"status"`
	RequestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		Page  *int `json:"Page" name:"Page"`
		Size  *int `json:"Size" name:"Size"`
		Total *int `json:"Total" name:"Total"`
		Bills []struct {
			AliasName               *string `json:"aliasName" name:"aliasName"`
			Email                   *string `json:"email" name:"email"`
			MembershipGroup         *string `json:"membershipGroup" name:"membershipGroup"`
			UserId                  *int    `json:"userId" name:"userId"`
			UserName                *string `json:"userName" name:"userName"`
			SellerCompanyName       *string `json:"sellerCompanyName" name:"sellerCompanyName"`
			BillMonth               *string `json:"billMonth" name:"billMonth"`
			CustomerBillMonth       *int    `json:"customerBillMonth" name:"customerBillMonth"`
			CurrencyCode            *string `json:"currencyCode" name:"currencyCode"`
			CurrencyInfo            *string `json:"currencyInfo" name:"currencyInfo"`
			ExchangeRate            *string `json:"exchangeRate" name:"exchangeRate"`
			BillDay                 *int    `json:"billDay" name:"billDay"`
			FinanceUnitName         *string `json:"financeUnitName" name:"financeUnitName"`
			PrePayCost              *string `json:"prePayCost" name:"prePayCost"`
			PrePayRefundAmount      *string `json:"prePayRefundAmount" name:"prePayRefundAmount"`
			PrePayBillsAmount       *string `json:"prePayBillsAmount" name:"prePayBillsAmount"`
			RealtimeCost            *string `json:"realtimeCost" name:"realtimeCost"`
			RealtimeAdjustAmount    *string `json:"realtimeAdjustAmount" name:"realtimeAdjustAmount"`
			RealtimeBillsAmount     *string `json:"realtimeBillsAmount" name:"realtimeBillsAmount"`
			PostPayCost             *string `json:"postPayCost" name:"postPayCost"`
			PostPayAdjustAmount     *string `json:"postPayAdjustAmount" name:"postPayAdjustAmount"`
			PostPayBillsAmount      *string `json:"postPayBillsAmount" name:"postPayBillsAmount"`
			BillFinalAmount         *string `json:"billFinalAmount" name:"billFinalAmount"`
			Cash                    *string `json:"cash" name:"cash"`
			Reward                  *string `json:"reward" name:"reward"`
			Vouchers                *string `json:"vouchers" name:"vouchers"`
			CloudTicketDenomination *string `json:"cloudTicketDenomination" name:"cloudTicketDenomination"`
		} `json:"Bills" name:"Bills"`
	} `json:"Data"`
	Error *string `json:"error" name:"error"`
	Ok    *bool   `json:"ok" name:"ok"`
}

func (r *QueryFinanceUnitConsumeOfMonthResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryFinanceUnitConsumeOfMonthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryUserConsumeRequest struct {
	*ksyunhttp.BaseRequest
	StartDay *string `json:"StartDay,omitempty" name:"StartDay"`
	EndDay   *string `json:"EndDay,omitempty" name:"EndDay"`
	Page     *int    `json:"Page,omitempty" name:"Page"`
	Size     *int    `json:"Size,omitempty" name:"Size"`
}

func (r *QueryUserConsumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUserConsumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryUserConsumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryUserConsumeResponse struct {
	*ksyunhttp.BaseResponse
	Status    *int    `json:"status" name:"status"`
	RequestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		Page  *int `json:"Page" name:"Page"`
		Size  *int `json:"Size" name:"Size"`
		Total *int `json:"Total" name:"Total"`
		Bills []struct {
			AliasName               *string `json:"aliasName" name:"aliasName"`
			Email                   *string `json:"email" name:"email"`
			MembershipGroup         *string `json:"membershipGroup" name:"membershipGroup"`
			UserId                  *int    `json:"userId" name:"userId"`
			UserName                *string `json:"userName" name:"userName"`
			SellerCompanyName       *string `json:"sellerCompanyName" name:"sellerCompanyName"`
			BillMonth               *string `json:"billMonth" name:"billMonth"`
			CustomerBillMonth       *int    `json:"customerBillMonth" name:"customerBillMonth"`
			CurrencyCode            *string `json:"currencyCode" name:"currencyCode"`
			CurrencyInfo            *string `json:"currencyInfo" name:"currencyInfo"`
			ExchangeRate            *string `json:"exchangeRate" name:"exchangeRate"`
			BillDay                 *int    `json:"billDay" name:"billDay"`
			PrePayCost              *string `json:"prePayCost" name:"prePayCost"`
			PrePayRefundAmount      *string `json:"prePayRefundAmount" name:"prePayRefundAmount"`
			PrePayBillsAmount       *string `json:"prePayBillsAmount" name:"prePayBillsAmount"`
			RealtimeCost            *string `json:"realtimeCost" name:"realtimeCost"`
			RealtimeAdjustAmount    *string `json:"realtimeAdjustAmount" name:"realtimeAdjustAmount"`
			RealtimeBillsAmount     *string `json:"realtimeBillsAmount" name:"realtimeBillsAmount"`
			PostPayCost             *string `json:"postPayCost" name:"postPayCost"`
			PostPayAdjustAmount     *string `json:"postPayAdjustAmount" name:"postPayAdjustAmount"`
			PostPayBillsAmount      *string `json:"postPayBillsAmount" name:"postPayBillsAmount"`
			BillFinalAmount         *string `json:"billFinalAmount" name:"billFinalAmount"`
			Cash                    *string `json:"cash" name:"cash"`
			Reward                  *string `json:"reward" name:"reward"`
			Vouchers                *string `json:"vouchers" name:"vouchers"`
			CloudTicketDenomination *string `json:"cloudTicketDenomination" name:"cloudTicketDenomination"`
		} `json:"Bills" name:"Bills"`
	} `json:"Data"`
	Error *string `json:"error" name:"error"`
	Ok    *bool   `json:"ok" name:"ok"`
}

func (r *QueryUserConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUserConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
