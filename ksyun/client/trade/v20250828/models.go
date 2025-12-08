package v20250828

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type QueryInstancesRequest struct {
	*ksyunhttp.BaseRequest
	AssociatedUserId     *int      `json:"associatedUserId,omitempty" name:"associatedUserId"`
	InstanceIds          []*string `json:"instanceIds,omitempty" name:"instanceIds"`
	Status               *int      `json:"status,omitempty" name:"status"`
	ProductGroup         *int      `json:"productGroup,omitempty" name:"productGroup"`
	RenewStrategy        *int      `json:"renewStrategy,omitempty" name:"renewStrategy"`
	BillingBeginTimeFrom *string   `json:"billingBeginTimeFrom,omitempty" name:"billingBeginTimeFrom"`
	BillingBeginTimeTo   *string   `json:"billingBeginTimeTo,omitempty" name:"billingBeginTimeTo"`
	BillingEndTimeFrom   *string   `json:"billingEndTimeFrom,omitempty" name:"billingEndTimeFrom"`
	BillingEndTimeTo     *string   `json:"billingEndTimeTo,omitempty" name:"billingEndTimeTo"`
	ServiceBeginTimeFrom *string   `json:"serviceBeginTimeFrom,omitempty" name:"serviceBeginTimeFrom"`
	ServiceBeginTimeTo   *string   `json:"serviceBeginTimeTo,omitempty" name:"serviceBeginTimeTo"`
	Page                 *int      `json:"page,omitempty" name:"page"`
	Size                 *int      `json:"size,omitempty" name:"size"`
}

func (r *QueryInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		RetResponse []struct {
			BillType          *int    `json:"billType" name:"billType"`
			BillTypeName      *string `json:"billTypeName" name:"billTypeName"`
			BillingBeginTime  *string `json:"billingBeginTime" name:"billingBeginTime"`
			BillingEndTime    *string `json:"billingEndTime" name:"billingEndTime"`
			CreateTime        *string `json:"createTime" name:"createTime"`
			IamProjectId      *int    `json:"iamProjectId" name:"iamProjectId"`
			InstanceId        *string `json:"instanceId" name:"instanceId"`
			InstanceType      *int    `json:"instanceType" name:"instanceType"`
			InstanceTypeName  *string `json:"instanceTypeName" name:"instanceTypeName"`
			PayTypeName       *string `json:"payTypeName" name:"payTypeName"`
			ProductGroup      *int    `json:"productGroup" name:"productGroup"`
			ProductType       *int    `json:"productType" name:"productType"`
			ProductTypeName   *string `json:"productTypeName" name:"productTypeName"`
			Region            *string `json:"region" name:"region"`
			RegionName        *string `json:"regionName" name:"regionName"`
			RenewDuration     *int    `json:"renewDuration" name:"renewDuration"`
			RenewStrategy     *int    `json:"renewStrategy" name:"renewStrategy"`
			RenewStrategyName *string `json:"renewStrategyName" name:"renewStrategyName"`
			Status            *int    `json:"status" name:"status"`
			StatusName        *string `json:"statusName" name:"statusName"`
			UserId            *int    `json:"userId" name:"userId"`
		} `json:"RetResponse" name:"RetResponse"`
		Total *int `json:"Total" name:"Total"`
	} `json:"Data"`
}

func (r *QueryInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
