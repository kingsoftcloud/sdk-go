package v20211201

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type GetRefreshOrPreloadTaskUrls struct {
	Url *string `json:"Url,omitempty" name:"Url"`
}

type GetRefreshOrPreloadTaskRequest struct {
	*ksyunhttp.BaseRequest
	StartTime  *string                        `json:"StartTime,omitempty" name:"StartTime"`
	EndTime    *string                        `json:"EndTime,omitempty" name:"EndTime"`
	TaskId     *string                        `json:"TaskId,omitempty" name:"TaskId"`
	DomainName *string                        `json:"DomainName,omitempty" name:"DomainName"`
	Urls       []*GetRefreshOrPreloadTaskUrls `json:"Urls,omitempty" name:"Urls"`
	Type       *string                        `json:"Type,omitempty" name:"Type"`
	SubType    *string                        `json:"SubType,omitempty" name:"SubType"`
	PageSize   *int64                         `json:"PageSize,omitempty" name:"PageSize"`
	PageNumber *int64                         `json:"PageNumber,omitempty" name:"PageNumber"`
}

func (r *GetRefreshOrPreloadTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRefreshOrPreloadTaskResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	Urls      []struct {
		Url *string `json:"Url" name:"Url"`
	} `json:"Urls"`
	PageSize   *int64 `json:"PageSize" name:"PageSize"`
	PageNumber *int64 `json:"PageNumber" name:"PageNumber"`
	TotalCount *int64 `json:"TotalCount" name:"TotalCount"`
	Datas      []struct {
		Type       *string  `json:"Type" name:"Type"`
		SubType    *string  `json:"SubType" name:"SubType"`
		Url        *string  `json:"Url" name:"Url"`
		Progress   *float64 `json:"Progress" name:"Progress"`
		Status     *string  `json:"Status" name:"Status"`
		TaskId     *string  `json:"TaskId" name:"TaskId"`
		CreateTime *string  `json:"CreateTime" name:"CreateTime"`
	} `json:"Datas"`
	Message *string `json:"Message" name:"Message"`
}

func (r *GetRefreshOrPreloadTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRefreshOrPreloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
