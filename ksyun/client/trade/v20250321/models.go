package v20250321
import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type QueryUnPayOrdersRequest struct {
	*ksyunhttp.BaseRequest
	Page *int `json:"page,omitempty" name:"page"`
	Size *int `json:"size,omitempty" name:"size"`
}

func (r *QueryUnPayOrdersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUnPayOrdersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryUnPayOrdersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryUnPayOrdersResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		BillName       *string `json:"BillName" name:"BillName"`
		CreateTime     *int    `json:"CreateTime" name:"CreateTime"`
		DiscountAmount *int    `json:"DiscountAmount" name:"DiscountAmount"`
		HasRefund      *int    `json:"HasRefund" name:"HasRefund"`
		OrderId        *string `json:"OrderId" name:"OrderId"`
		OrderList      []struct {
			BillType         *int    `json:"billType" name:"billType"`
			IamProjectId     *int    `json:"iamProjectId" name:"iamProjectId"`
			Num              *int    `json:"num" name:"num"`
			OldPrice         *int    `json:"oldPrice" name:"oldPrice"`
			OrderId          *string `json:"orderId" name:"orderId"`
			Price            *int    `json:"price" name:"price"`
			ProductGroup     *int    `json:"productGroup" name:"productGroup"`
			ProductGroupName *string `json:"productGroupName" name:"productGroupName"`
			ProductId        *string `json:"productId" name:"productId"`
			ProductItems     []struct {
				Key *string `json:"Key" name:"Key"`
				Value *string `json:"Value" name:"Value"`
			} `json:"ProductItems"`
			ProductType     *int    `json:"productType" name:"productType"`
			ProductTypeName *string `json:"productTypeName" name:"productTypeName"`
			PromotionItem   struct {
				Name *string `json:"Name" name:"Name"`
				Type *string `json:"Type" name:"Type"`
			} `json:"PromotionItem"`
			ServiceBeginTime *int `json:"serviceBeginTime" name:"serviceBeginTime"`
			ServiceEndTime   *int `json:"serviceEndTime" name:"serviceEndTime"`
			Status           *int `json:"status" name:"status"`
		} `json:"OrderList" name:"OrderList"`
		OrderName    *string `json:"OrderName" name:"OrderName"`
		OrderType    *int    `json:"OrderType" name:"OrderType"`
		OrderUse     *int    `json:"OrderUse" name:"OrderUse"`
		PayableMoney *int    `json:"PayableMoney" name:"PayableMoney"`
		RealMoney    *int    `json:"RealMoney" name:"RealMoney"`
		Status       *int    `json:"Status" name:"Status"`
		TotalMoney   *int    `json:"TotalMoney" name:"TotalMoney"`
		UpdateTime   *int    `json:"UpdateTime" name:"UpdateTime"`
		UserId       *int    `json:"UserId" name:"UserId"`
	} `json:"Data"`
	Total *int `json:"total" name:"total"`
}

func (r *QueryUnPayOrdersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryUnPayOrdersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type QueryOrderInfoRequest struct {
	*ksyunhttp.BaseRequest
	OrderId *string `json:"orderId,omitempty" name:"orderId"`
}

func (r *QueryOrderInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryOrderInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "QueryOrderInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type QueryOrderInfoResponse struct {
	*ksyunhttp.BaseResponse
	BillName       *string `json:"billName" name:"billName"`
	CreateTime     *int    `json:"createTime" name:"createTime"`
	DiscountAmount *int    `json:"discountAmount" name:"discountAmount"`
	HasRefund      *int    `json:"hasRefund" name:"hasRefund"`
	OrderId        *string `json:"orderId" name:"orderId"`
	OrderList      []struct {
		BillType     *int    `json:"BillType" name:"BillType"`
		IamProjectId *int    `json:"IamProjectId" name:"IamProjectId"`
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		Num          *int    `json:"Num" name:"Num"`
		OrderId      *string `json:"OrderId" name:"OrderId"`
		Price        *int    `json:"Price" name:"Price"`
		ProductGroup *int    `json:"ProductGroup" name:"ProductGroup"`
		ProductGroupName *string `json:"ProductGroupName" name:"ProductGroupName"`
		ProductId    *string `json:"ProductId" name:"ProductId"`
		ProductItems []struct {
			Key   *string `json:"key" name:"key"`
			Value *string `json:"value" name:"value"`
		} `json:"ProductItems" name:"ProductItems"`
		ProductType   *int `json:"ProductType" name:"ProductType"`
		ProductTypeName *string `json:"ProductTypeName" name:"ProductTypeName"`
		PromotionItem struct {
			Name *string `json:"name" name:"name"`
			Type *string `json:"type" name:"type"`
		} `json:"PromotionItem" name:"PromotionItem"`
		ServiceBeginTime *int    `json:"ServiceBeginTime" name:"ServiceBeginTime"`
		ServiceEndTime   *int    `json:"ServiceEndTime" name:"ServiceEndTime"`
		Status           *int    `json:"Status" name:"Status"`
		SubOrderId       *string `json:"SubOrderId" name:"SubOrderId"`
	} `json:"OrderList"`
	OrderName    *string `json:"orderName" name:"orderName"`
	OrderType    *int    `json:"orderType" name:"orderType"`
	OrderUse     *int    `json:"orderUse" name:"orderUse"`
	PayTime      *int    `json:"payTime" name:"payTime"`
	PayableMoney *int    `json:"payableMoney" name:"payableMoney"`
	RealMoney    *int    `json:"realMoney" name:"realMoney"`
	Status       *int    `json:"status" name:"status"`
	TotalMoney   *int    `json:"totalMoney" name:"totalMoney"`
	UpdateTime   *int    `json:"updateTime" name:"updateTime"`
	UserId       *int    `json:"userId" name:"userId"`
}

func (r *QueryOrderInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryOrderInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CancelOrderRequest struct {
	*ksyunhttp.BaseRequest
	OrderId *string `json:"orderId,omitempty" name:"orderId"`
}

func (r *CancelOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CancelOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CancelOrderResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Success   *bool   `json:"success" name:"success"`
}

func (r *CancelOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type LaunchPayOrderRequest struct {
	*ksyunhttp.BaseRequest
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
}

func (r *LaunchPayOrderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LaunchPayOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "LaunchPayOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type LaunchPayOrderResponse struct {
	*ksyunhttp.BaseResponse
	OrderId    *string `json:"OrderId" name:"OrderId"`
	RealMoney  *int    `json:"RealMoney" name:"RealMoney"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	SubOrderIds []*string `json:"SubOrderIds" name:"SubOrderIds"`
	TotalMoney *int    `json:"TotalMoney" name:"TotalMoney"`
}

func (r *LaunchPayOrderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *LaunchPayOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

