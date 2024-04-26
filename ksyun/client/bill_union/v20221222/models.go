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
		Page  *int `json:"page"`
		Size  *int `json:"size"`
		Total *int `json:"total"`
		Bills []struct {
			AliasName               *string `json:"aliasName"`
			Email                   *string `json:"email"`
			MembershipGroup         *string `json:"membershipGroup"`
			UserId                  *int    `json:"userId"`
			UserName                *string `json:"userName"`
			SellerCompanyName       *string `json:"sellerCompanyName"`
			BillMonth               *int    `json:"billMonth"`
			CustomerBillMonth       *int    `json:"customerBillMonth"`
			CurrencyCode            *string `json:"currencyCode"`
			CurrencyInfo            *string `json:"currencyInfo"`
			ExchangeRate            *string `json:"exchangeRate"`
			BillDay                 *int    `json:"billDay"`
			Id                      *string `json:"id"`
			FinanceUnitName         *string `json:"financeUnitName"`
			BillStartTime           *string `json:"billStartTime"`
			BillEndTime             *string `json:"billEndTime"`
			InstanceId              *string `json:"instanceId"`
			InstanceName            *string `json:"instanceName"`
			ProductTypeId           *int    `json:"productTypeId"`
			ProductTypeName         *string `json:"productTypeName"`
			ProductGroupId          *int    `json:"productGroupId"`
			ProductGroupName        *string `json:"productGroupName"`
			PayType                 *int    `json:"payType"`
			BillRealAmount          *string `json:"billRealAmount"`
			OriginalAmount          *string `json:"originalAmount"`
			RegionCode              *string `json:"regionCode"`
			ProjectId               *int    `json:"projectId"`
			ProjectName             *string `json:"projectName"`
			RegionName              *string `json:"regionName"`
			BillType                *int    `json:"billType"`
			BillTypeName            *string `json:"billTypeName"`
			PayTypeName             *string `json:"payTypeName"`
			BillDetailType          *int    `json:"billDetailType"`
			BillDetailTypeName      *string `json:"billDetailTypeName"`
			Duration                *string `json:"duration"`
			DurationNumber          *int    `json:"durationNumber"`
			RuleRemark              *string `json:"ruleRemark"`
			AvailabilityZone        *string `json:"availabilityZone"`
			Discount                *string `json:"discount"`
			Cash                    *string `json:"cash"`
			Reward                  *string `json:"reward"`
			Vouchers                *string `json:"vouchers"`
			CloudTicketDenomination *string `json:"cloudTicketDenomination"`
			CloudTicketCost         *string `json:"cloudTicketCost"`
			ServiceBegionTime       *string `json:"serviceBegionTime"`
			ConfigInfo              []struct {
			} `json:"configInfo"`
			PriceFactorInfo []struct {
			} `json:"priceFactorInfo"`
			ExtraInfo []struct {
			} `json:"extraInfo"`
			ResourceDeductionInfo []struct {
			} `json:"resourceDeductionInfo"`
			TagInfo []struct {
			} `json:"tagInfo"`
		} `json:"Bills"`
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
	status    *int    `json:"status" name:"status"`
	requestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		Page  *int `json:"Page"`
		Size  *int `json:"Size"`
		Total *int `json:"Total"`
		Bills []struct {
			aliasName               *string `json:"aliasName"`
			email                   *string `json:"email"`
			membershipGroup         *string `json:"membershipGroup"`
			userId                  *int    `json:"userId"`
			userName                *string `json:"userName"`
			sellerCompanyName       *string `json:"sellerCompanyName"`
			billMonth               *string `json:"billMonth"`
			customerBillMonth       *int    `json:"customerBillMonth"`
			currencyCode            *string `json:"currencyCode"`
			currencyInfo            *string `json:"currencyInfo"`
			exchangeRate            *string `json:"exchangeRate"`
			billDay                 *int    `json:"billDay"`
			projectId               *int    `json:"projectId"`
			projectName             *string `json:"projectName"`
			prePayCost              *string `json:"prePayCost"`
			prePayRefundAmount      *string `json:"prePayRefundAmount"`
			prePayBillsAmount       *string `json:"prePayBillsAmount"`
			realtimeCost            *string `json:"realtimeCost"`
			realtimeAdjustAmount    *string `json:"realtimeAdjustAmount"`
			realtimeBillsAmount     *string `json:"realtimeBillsAmount"`
			postPayCost             *string `json:"postPayCost"`
			postPayAdjustAmount     *string `json:"postPayAdjustAmount"`
			postPayBillsAmount      *string `json:"postPayBillsAmount"`
			billFinalAmount         *string `json:"billFinalAmount"`
			cash                    *string `json:"cash"`
			reward                  *string `json:"reward"`
			vouchers                *string `json:"vouchers"`
			cloudTicketDenomination *string `json:"cloudTicketDenomination"`
		} `json:"Bills"`
	} `json:"Data"`
	error *string `json:"error" name:"error"`
	ok    *bool   `json:"ok" name:"ok"`
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
	status    *int    `json:"status" name:"status"`
	requestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		Page  *int `json:"Page"`
		Size  *int `json:"Size"`
		Total *int `json:"Total"`
		Bills []struct {
			aliasName               *string `json:"aliasName"`
			email                   *string `json:"email"`
			membershipGroup         *string `json:"membershipGroup"`
			userId                  *int    `json:"userId"`
			userName                *string `json:"userName"`
			sellerCompanyName       *string `json:"sellerCompanyName"`
			billMonth               *string `json:"billMonth"`
			customerBillMonth       *int    `json:"customerBillMonth"`
			currencyCode            *string `json:"currencyCode"`
			currencyInfo            *string `json:"currencyInfo"`
			exchangeRate            *string `json:"exchangeRate"`
			billDay                 *int    `json:"billDay"`
			productGroupId          *int    `json:"productGroupId"`
			productGroupName        *string `json:"productGroupName"`
			prePayCost              *string `json:"prePayCost"`
			prePayRefundAmount      *string `json:"prePayRefundAmount"`
			prePayBillsAmount       *string `json:"prePayBillsAmount"`
			realtimeCost            *string `json:"realtimeCost"`
			realtimeAdjustAmount    *string `json:"realtimeAdjustAmount"`
			realtimeBillsAmount     *string `json:"realtimeBillsAmount"`
			postPayCost             *string `json:"postPayCost"`
			postPayAdjustAmount     *string `json:"postPayAdjustAmount"`
			postPayBillsAmount      *string `json:"postPayBillsAmount"`
			billFinalAmount         *string `json:"billFinalAmount"`
			cash                    *string `json:"cash"`
			reward                  *string `json:"reward"`
			vouchers                *string `json:"vouchers"`
			cloudTicketDenomination *string `json:"cloudTicketDenomination"`
		} `json:"Bills"`
	} `json:"Data"`
	error *string `json:"error" name:"error"`
	ok    *bool   `json:"ok" name:"ok"`
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
	status    *int    `json:"status" name:"status"`
	requestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		Page  *int `json:"Page"`
		Size  *int `json:"Size"`
		Total *int `json:"Total"`
		Bills []struct {
			aliasName               *string `json:"aliasName"`
			email                   *string `json:"email"`
			membershipGroup         *string `json:"membershipGroup"`
			userId                  *int    `json:"userId"`
			userName                *string `json:"userName"`
			sellerCompanyName       *string `json:"sellerCompanyName"`
			billMonth               *string `json:"billMonth"`
			customerBillMonth       *int    `json:"customerBillMonth"`
			currencyCode            *string `json:"currencyCode"`
			currencyInfo            *string `json:"currencyInfo"`
			exchangeRate            *string `json:"exchangeRate"`
			billDay                 *int    `json:"billDay"`
			financeUnitName         *string `json:"financeUnitName"`
			prePayCost              *string `json:"prePayCost"`
			prePayRefundAmount      *string `json:"prePayRefundAmount"`
			prePayBillsAmount       *string `json:"prePayBillsAmount"`
			realtimeCost            *string `json:"realtimeCost"`
			realtimeAdjustAmount    *string `json:"realtimeAdjustAmount"`
			realtimeBillsAmount     *string `json:"realtimeBillsAmount"`
			postPayCost             *string `json:"postPayCost"`
			postPayAdjustAmount     *string `json:"postPayAdjustAmount"`
			postPayBillsAmount      *string `json:"postPayBillsAmount"`
			billFinalAmount         *string `json:"billFinalAmount"`
			cash                    *string `json:"cash"`
			reward                  *string `json:"reward"`
			vouchers                *string `json:"vouchers"`
			cloudTicketDenomination *string `json:"cloudTicketDenomination"`
		} `json:"Bills"`
	} `json:"Data"`
	error *string `json:"error" name:"error"`
	ok    *bool   `json:"ok" name:"ok"`
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
	status    *int    `json:"status" name:"status"`
	requestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		Page  *int `json:"Page"`
		Size  *int `json:"Size"`
		Total *int `json:"Total"`
		Bills []struct {
			aliasName               *string `json:"aliasName"`
			email                   *string `json:"email"`
			membershipGroup         *string `json:"membershipGroup"`
			userId                  *int    `json:"userId"`
			userName                *string `json:"userName"`
			sellerCompanyName       *string `json:"sellerCompanyName"`
			billMonth               *string `json:"billMonth"`
			customerBillMonth       *int    `json:"customerBillMonth"`
			currencyCode            *string `json:"currencyCode"`
			currencyInfo            *string `json:"currencyInfo"`
			exchangeRate            *string `json:"exchangeRate"`
			billDay                 *int    `json:"billDay"`
			financeUnitName         *string `json:"financeUnitName"`
			prePayCost              *string `json:"prePayCost"`
			prePayRefundAmount      *string `json:"prePayRefundAmount"`
			prePayBillsAmount       *string `json:"prePayBillsAmount"`
			realtimeCost            *string `json:"realtimeCost"`
			realtimeAdjustAmount    *string `json:"realtimeAdjustAmount"`
			realtimeBillsAmount     *string `json:"realtimeBillsAmount"`
			postPayCost             *string `json:"postPayCost"`
			postPayAdjustAmount     *string `json:"postPayAdjustAmount"`
			postPayBillsAmount      *string `json:"postPayBillsAmount"`
			billFinalAmount         *string `json:"billFinalAmount"`
			cash                    *string `json:"cash"`
			reward                  *string `json:"reward"`
			vouchers                *string `json:"vouchers"`
			cloudTicketDenomination *string `json:"cloudTicketDenomination"`
		} `json:"Bills"`
	} `json:"Data"`
	error *string `json:"error" name:"error"`
	ok    *bool   `json:"ok" name:"ok"`
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
	status    *int    `json:"status" name:"status"`
	requestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		Page  *int `json:"Page"`
		Size  *int `json:"Size"`
		Total *int `json:"Total"`
		Bills []struct {
			aliasName               *string `json:"aliasName"`
			email                   *string `json:"email"`
			membershipGroup         *string `json:"membershipGroup"`
			userId                  *int    `json:"userId"`
			userName                *string `json:"userName"`
			sellerCompanyName       *string `json:"sellerCompanyName"`
			billMonth               *string `json:"billMonth"`
			customerBillMonth       *int    `json:"customerBillMonth"`
			currencyCode            *string `json:"currencyCode"`
			currencyInfo            *string `json:"currencyInfo"`
			exchangeRate            *string `json:"exchangeRate"`
			billDay                 *int    `json:"billDay"`
			prePayCost              *string `json:"prePayCost"`
			prePayRefundAmount      *string `json:"prePayRefundAmount"`
			prePayBillsAmount       *string `json:"prePayBillsAmount"`
			realtimeCost            *string `json:"realtimeCost"`
			realtimeAdjustAmount    *string `json:"realtimeAdjustAmount"`
			realtimeBillsAmount     *string `json:"realtimeBillsAmount"`
			postPayCost             *string `json:"postPayCost"`
			postPayAdjustAmount     *string `json:"postPayAdjustAmount"`
			postPayBillsAmount      *string `json:"postPayBillsAmount"`
			billFinalAmount         *string `json:"billFinalAmount"`
			cash                    *string `json:"cash"`
			reward                  *string `json:"reward"`
			vouchers                *string `json:"vouchers"`
			cloudTicketDenomination *string `json:"cloudTicketDenomination"`
		} `json:"Bills"`
	} `json:"Data"`
	error *string `json:"error" name:"error"`
	ok    *bool   `json:"ok" name:"ok"`
}

func (r *QueryUserConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUserConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
