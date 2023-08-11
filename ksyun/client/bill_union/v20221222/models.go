package v20221222
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type QueryInstanceConsumeRequest struct {
    *ksyunhttp.BaseRequest
    StartDay *string `json:"StartDay,omitempty" name:"StartDay"`
    EndDay *string `json:"EndDay,omitempty" name:"EndDay"`
    ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
    Page *int `json:"Page,omitempty" name:"Page"`
    Size *int `json:"Size,omitempty" name:"Size"`
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
    status *int `json:"status" name:"status"`
    requestId *string `json:"requestId" name:"requestId"`
	Data struct {
		Page *int `json:"Page"`
		Size *int `json:"Size"`
		Total *int `json:"Total"`
		Bills []struct {
					aliasName *string `json:"aliasName"`
					email *string `json:"email"`
					membershipGroup *string `json:"membershipGroup"`
					userId *int `json:"userId"`
					userName *string `json:"userName"`
					sellerCompanyName *string `json:"sellerCompanyName"`
					billMonth *int `json:"billMonth"`
					customerBillMonth *int `json:"customerBillMonth"`
					currencyCode *string `json:"currencyCode"`
					currencyInfo *string `json:"currencyInfo"`
					exchangeRate *string `json:"exchangeRate"`
					billDay *int `json:"billDay"`
					id *string `json:"id"`
					financeUnitName *string `json:"financeUnitName"`
					billStartTime *string `json:"billStartTime"`
					billEndTime *string `json:"billEndTime"`
					instanceId *string `json:"instanceId"`
					instanceName *string `json:"instanceName"`
					productTypeId *int `json:"productTypeId"`
					productTypeName *string `json:"productTypeName"`
					productGroupId *int `json:"productGroupId"`
					productGroupName *string `json:"productGroupName"`
					payType *int `json:"payType"`
					billRealAmount *string `json:"billRealAmount"`
					originalAmount *string `json:"originalAmount"`
					regionCode *string `json:"regionCode"`
					projectId *int `json:"projectId"`
					projectName *string `json:"projectName"`
					regionName *string `json:"regionName"`
					billType *int `json:"billType"`
					billTypeName *string `json:"billTypeName"`
					payTypeName *string `json:"payTypeName"`
					billDetailType *int `json:"billDetailType"`
					billDetailTypeName *string `json:"billDetailTypeName"`
					duration *string `json:"duration"`
					durationNumber *int `json:"durationNumber"`
					ruleRemark *string `json:"ruleRemark"`
					availabilityZone *string `json:"availabilityZone"`
					discount *string `json:"discount"`
					cash *string `json:"cash"`
					reward *string `json:"reward"`
					vouchers *string `json:"vouchers"`
					cloudTicketDenomination *string `json:"cloudTicketDenomination"`
					cloudTicketCost *string `json:"cloudTicketCost"`
					serviceBegionTime *string `json:"serviceBegionTime"`
				ConfigInfo []struct {
				} `json:"ConfigInfo"`
				PriceFactorInfo []struct {
				} `json:"PriceFactorInfo"`
				ExtraInfo []struct {
				} `json:"ExtraInfo"`
				ResourceDeductionInfo []struct {
				} `json:"ResourceDeductionInfo"`
				TagInfo []struct {
				} `json:"TagInfo"`
			} `json:"Bills"`
		} `json:"Data"`
    error *string `json:"error" name:"error"`
    ok *bool `json:"ok" name:"ok"`
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
    EndDay *string `json:"EndDay,omitempty" name:"EndDay"`
    Page *int `json:"Page,omitempty" name:"Page"`
    Size *int `json:"Size,omitempty" name:"Size"`
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
    status *int `json:"status" name:"status"`
    requestId *string `json:"requestId" name:"requestId"`
	Data struct {
		Page *int `json:"Page"`
		Size *int `json:"Size"`
		Total *int `json:"Total"`
		Bills []struct {
					aliasName *string `json:"aliasName"`
					email *string `json:"email"`
					membershipGroup *string `json:"membershipGroup"`
					userId *int `json:"userId"`
					userName *string `json:"userName"`
					sellerCompanyName *string `json:"sellerCompanyName"`
					billMonth *string `json:"billMonth"`
					customerBillMonth *int `json:"customerBillMonth"`
					currencyCode *string `json:"currencyCode"`
					currencyInfo *string `json:"currencyInfo"`
					exchangeRate *string `json:"exchangeRate"`
					billDay *int `json:"billDay"`
					projectId *int `json:"projectId"`
					projectName *string `json:"projectName"`
					prePayCost *string `json:"prePayCost"`
					prePayRefundAmount *string `json:"prePayRefundAmount"`
					prePayBillsAmount *string `json:"prePayBillsAmount"`
					realtimeCost *string `json:"realtimeCost"`
					realtimeAdjustAmount *string `json:"realtimeAdjustAmount"`
					realtimeBillsAmount *string `json:"realtimeBillsAmount"`
					postPayCost *string `json:"postPayCost"`
					postPayAdjustAmount *string `json:"postPayAdjustAmount"`
					postPayBillsAmount *string `json:"postPayBillsAmount"`
					billFinalAmount *string `json:"billFinalAmount"`
					cash *string `json:"cash"`
					reward *string `json:"reward"`
					vouchers *string `json:"vouchers"`
					cloudTicketDenomination *string `json:"cloudTicketDenomination"`
			} `json:"Bills"`
		} `json:"Data"`
    error *string `json:"error" name:"error"`
    ok *bool `json:"ok" name:"ok"`
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
    EndDay *string `json:"EndDay,omitempty" name:"EndDay"`
    Page *int `json:"Page,omitempty" name:"Page"`
    Size *int `json:"Size,omitempty" name:"Size"`
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
    status *int `json:"status" name:"status"`
    requestId *string `json:"requestId" name:"requestId"`
	Data struct {
		Page *int `json:"Page"`
		Size *int `json:"Size"`
		Total *int `json:"Total"`
		Bills []struct {
					aliasName *string `json:"aliasName"`
					email *string `json:"email"`
					membershipGroup *string `json:"membershipGroup"`
					userId *int `json:"userId"`
					userName *string `json:"userName"`
					sellerCompanyName *string `json:"sellerCompanyName"`
					billMonth *string `json:"billMonth"`
					customerBillMonth *int `json:"customerBillMonth"`
					currencyCode *string `json:"currencyCode"`
					currencyInfo *string `json:"currencyInfo"`
					exchangeRate *string `json:"exchangeRate"`
					billDay *int `json:"billDay"`
					productGroupId *int `json:"productGroupId"`
					productGroupName *string `json:"productGroupName"`
					prePayCost *string `json:"prePayCost"`
					prePayRefundAmount *string `json:"prePayRefundAmount"`
					prePayBillsAmount *string `json:"prePayBillsAmount"`
					realtimeCost *string `json:"realtimeCost"`
					realtimeAdjustAmount *string `json:"realtimeAdjustAmount"`
					realtimeBillsAmount *string `json:"realtimeBillsAmount"`
					postPayCost *string `json:"postPayCost"`
					postPayAdjustAmount *string `json:"postPayAdjustAmount"`
					postPayBillsAmount *string `json:"postPayBillsAmount"`
					billFinalAmount *string `json:"billFinalAmount"`
					cash *string `json:"cash"`
					reward *string `json:"reward"`
					vouchers *string `json:"vouchers"`
					cloudTicketDenomination *string `json:"cloudTicketDenomination"`
			} `json:"Bills"`
		} `json:"Data"`
    error *string `json:"error" name:"error"`
    ok *bool `json:"ok" name:"ok"`
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
    EndDay *string `json:"EndDay,omitempty" name:"EndDay"`
    Page *int `json:"Page,omitempty" name:"Page"`
    Size *int `json:"Size,omitempty" name:"Size"`
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
    status *int `json:"status" name:"status"`
    requestId *string `json:"requestId" name:"requestId"`
	Data struct {
		Page *int `json:"Page"`
		Size *int `json:"Size"`
		Total *int `json:"Total"`
		Bills []struct {
					aliasName *string `json:"aliasName"`
					email *string `json:"email"`
					membershipGroup *string `json:"membershipGroup"`
					userId *int `json:"userId"`
					userName *string `json:"userName"`
					sellerCompanyName *string `json:"sellerCompanyName"`
					billMonth *string `json:"billMonth"`
					customerBillMonth *int `json:"customerBillMonth"`
					currencyCode *string `json:"currencyCode"`
					currencyInfo *string `json:"currencyInfo"`
					exchangeRate *string `json:"exchangeRate"`
					billDay *int `json:"billDay"`
					financeUnitName *string `json:"financeUnitName"`
					prePayCost *string `json:"prePayCost"`
					prePayRefundAmount *string `json:"prePayRefundAmount"`
					prePayBillsAmount *string `json:"prePayBillsAmount"`
					realtimeCost *string `json:"realtimeCost"`
					realtimeAdjustAmount *string `json:"realtimeAdjustAmount"`
					realtimeBillsAmount *string `json:"realtimeBillsAmount"`
					postPayCost *string `json:"postPayCost"`
					postPayAdjustAmount *string `json:"postPayAdjustAmount"`
					postPayBillsAmount *string `json:"postPayBillsAmount"`
					billFinalAmount *string `json:"billFinalAmount"`
					cash *string `json:"cash"`
					reward *string `json:"reward"`
					vouchers *string `json:"vouchers"`
					cloudTicketDenomination *string `json:"cloudTicketDenomination"`
			} `json:"Bills"`
		} `json:"Data"`
    error *string `json:"error" name:"error"`
    ok *bool `json:"ok" name:"ok"`
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
    Page *int `json:"Page,omitempty" name:"Page"`
    Size *int `json:"Size,omitempty" name:"Size"`
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
    status *int `json:"status" name:"status"`
    requestId *string `json:"requestId" name:"requestId"`
	Data struct {
		Page *int `json:"Page"`
		Size *int `json:"Size"`
		Total *int `json:"Total"`
		Bills []struct {
					aliasName *string `json:"aliasName"`
					email *string `json:"email"`
					membershipGroup *string `json:"membershipGroup"`
					userId *int `json:"userId"`
					userName *string `json:"userName"`
					sellerCompanyName *string `json:"sellerCompanyName"`
					billMonth *string `json:"billMonth"`
					customerBillMonth *int `json:"customerBillMonth"`
					currencyCode *string `json:"currencyCode"`
					currencyInfo *string `json:"currencyInfo"`
					exchangeRate *string `json:"exchangeRate"`
					billDay *int `json:"billDay"`
					financeUnitName *string `json:"financeUnitName"`
					prePayCost *string `json:"prePayCost"`
					prePayRefundAmount *string `json:"prePayRefundAmount"`
					prePayBillsAmount *string `json:"prePayBillsAmount"`
					realtimeCost *string `json:"realtimeCost"`
					realtimeAdjustAmount *string `json:"realtimeAdjustAmount"`
					realtimeBillsAmount *string `json:"realtimeBillsAmount"`
					postPayCost *string `json:"postPayCost"`
					postPayAdjustAmount *string `json:"postPayAdjustAmount"`
					postPayBillsAmount *string `json:"postPayBillsAmount"`
					billFinalAmount *string `json:"billFinalAmount"`
					cash *string `json:"cash"`
					reward *string `json:"reward"`
					vouchers *string `json:"vouchers"`
					cloudTicketDenomination *string `json:"cloudTicketDenomination"`
			} `json:"Bills"`
		} `json:"Data"`
    error *string `json:"error" name:"error"`
    ok *bool `json:"ok" name:"ok"`
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
    EndDay *string `json:"EndDay,omitempty" name:"EndDay"`
    Page *int `json:"Page,omitempty" name:"Page"`
    Size *int `json:"Size,omitempty" name:"Size"`
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
    status *int `json:"status" name:"status"`
    requestId *string `json:"requestId" name:"requestId"`
	Data struct {
		Page *int `json:"Page"`
		Size *int `json:"Size"`
		Total *int `json:"Total"`
		Bills []struct {
					aliasName *string `json:"aliasName"`
					email *string `json:"email"`
					membershipGroup *string `json:"membershipGroup"`
					userId *int `json:"userId"`
					userName *string `json:"userName"`
					sellerCompanyName *string `json:"sellerCompanyName"`
					billMonth *string `json:"billMonth"`
					customerBillMonth *int `json:"customerBillMonth"`
					currencyCode *string `json:"currencyCode"`
					currencyInfo *string `json:"currencyInfo"`
					exchangeRate *string `json:"exchangeRate"`
					billDay *int `json:"billDay"`
					prePayCost *string `json:"prePayCost"`
					prePayRefundAmount *string `json:"prePayRefundAmount"`
					prePayBillsAmount *string `json:"prePayBillsAmount"`
					realtimeCost *string `json:"realtimeCost"`
					realtimeAdjustAmount *string `json:"realtimeAdjustAmount"`
					realtimeBillsAmount *string `json:"realtimeBillsAmount"`
					postPayCost *string `json:"postPayCost"`
					postPayAdjustAmount *string `json:"postPayAdjustAmount"`
					postPayBillsAmount *string `json:"postPayBillsAmount"`
					billFinalAmount *string `json:"billFinalAmount"`
					cash *string `json:"cash"`
					reward *string `json:"reward"`
					vouchers *string `json:"vouchers"`
					cloudTicketDenomination *string `json:"cloudTicketDenomination"`
			} `json:"Bills"`
		} `json:"Data"`
    error *string `json:"error" name:"error"`
    ok *bool `json:"ok" name:"ok"`
}

func (r *QueryUserConsumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryUserConsumeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

