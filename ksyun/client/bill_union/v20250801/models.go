package v20250801
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type QueryItemBillsRequest struct {
	*ksyunhttp.BaseRequest
	CustomerBillMonth *int    `json:"CustomerBillMonth,omitempty" name:"CustomerBillMonth"`
	ProductGroupCode  *string `json:"ProductGroupCode,omitempty" name:"ProductGroupCode"`
	PayType           *int    `json:"PayType,omitempty" name:"PayType"`
	InstanceId        *string `json:"InstanceId,omitempty" name:"InstanceId"`
	SubAccountId      *int    `json:"SubAccountId,omitempty" name:"SubAccountId"`
	Size              *int    `json:"Size,omitempty" name:"Size"`
	LastSortValue     *string `json:"LastSortValue,omitempty" name:"LastSortValue"`
}

func (r *QueryItemBillsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryItemBillsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
	} `json:"Data"`
}

func (r *QueryItemBillsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryItemBillsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type QueryProductTypesRequest struct {
	*ksyunhttp.BaseRequest
	ProductGroupCode *string `json:"ProductGroupCode,omitempty" name:"ProductGroupCode"`
}

func (r *QueryProductTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryProductTypesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		ProductGroupId *int `json:"ProductGroupId" name:"ProductGroupId"`
		ProductGroupName *string `json:"ProductGroupName" name:"ProductGroupName"`
		ProductTypeId  *int `json:"ProductTypeId" name:"ProductTypeId"`
		ProductTypeName *string `json:"ProductTypeName" name:"ProductTypeName"`
		ProductGroupCode *string `json:"ProductGroupCode" name:"ProductGroupCode"`
		ProductTypeCode *string `json:"ProductTypeCode" name:"ProductTypeCode"`
	} `json:"Data"`
}

func (r *QueryProductTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryProductTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

