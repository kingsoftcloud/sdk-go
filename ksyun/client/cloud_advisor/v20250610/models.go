package v20250610
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type GetReportRequest struct {
	*ksyunhttp.BaseRequest
	TaskIDs   []*string `json:"taskIDs,omitempty" name:"taskIDs"`
	StartTime *string   `json:"startTime,omitempty" name:"startTime"`
	EndTime   *string   `json:"endTime,omitempty" name:"endTime"`
}

func (r *GetReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetReportResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"code" name:"code"`
	Msg       *string `json:"msg" name:"msg"`
	RequestId *string `json:"requestId" name:"requestId"`
	Data      []struct {
		TaskId     *string `json:"TaskId" name:"TaskId"`
		FilePdfUrl *string `json:"FilePdfUrl" name:"FilePdfUrl"`
		FileExcelUrl *string `json:"FileExcelUrl" name:"FileExcelUrl"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		Status     *int    `json:"Status" name:"Status"`
		StatusName *string `json:"StatusName" name:"StatusName"`
	} `json:"Data"`
}

func (r *GetReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateTaskRequest struct {
	*ksyunhttp.BaseRequest
	InspectionItemIDs   []*int `json:"inspectionItemIDs,omitempty" name:"inspectionItemIDs"`
	ProductGroupIds     []*int `json:"productGroupIds,omitempty" name:"productGroupIds"`
	InspectionItemTypes []*int `json:"inspectionItemTypes,omitempty" name:"inspectionItemTypes"`
}

func (r *CreateTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateTaskResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"code" name:"code"`
	Msg       *string `json:"msg" name:"msg"`
	RequestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		TaskId *string `json:"TaskId" name:"TaskId"`
	} `json:"Data"`
}

func (r *CreateTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ListInspectionItemRequest struct {
	*ksyunhttp.BaseRequest
	PageNum  *int `json:"pageNum,omitempty" name:"pageNum"`
	PageSize *int `json:"pageSize,omitempty" name:"pageSize"`
}

func (r *ListInspectionItemRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListInspectionItemResponse struct {
	*ksyunhttp.BaseResponse
	Code      *int    `json:"code" name:"code"`
	Msg       *string `json:"msg" name:"msg"`
	RequestId *string `json:"requestId" name:"requestId"`
	Data      struct {
		Items []struct {
			CustomerId         *int    `json:"customerId" name:"customerId"`
			InspectionItemId   *int    `json:"inspectionItemId" name:"inspectionItemId"`
			InspectionItemName *string `json:"inspectionItemName" name:"inspectionItemName"`
			ProductGroupId     *int    `json:"productGroupId" name:"productGroupId"`
			ProductGroupName   *string `json:"productGroupName" name:"productGroupName"`
			TypeField          *int    `json:"type" name:"type"`
			TypeName           *string `json:"typeName" name:"typeName"`
			Status             *int    `json:"status" name:"status"`
			StatusName         *string `json:"StatusName" name:"StatusName"`
			Remark             *string `json:"remark" name:"remark"`
			Id                 *int    `json:"id" name:"id"`
			RiskDetails        *string `json:"riskDetails" name:"riskDetails"`
		} `json:"Items" name:"Items"`
		Total *int `json:"Total" name:"Total"`
		Page *int `json:"Page" name:"Page"`
		Size *int `json:"Size" name:"Size"`
	} `json:"Data"`
}

func (r *ListInspectionItemResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListInspectionItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

