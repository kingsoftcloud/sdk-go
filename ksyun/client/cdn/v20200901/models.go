package v20200901

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type CreateUserUsageDataExportTaskRequest struct {
	*ksyunhttp.BaseRequest
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" name:"EndTime"`
	CdnType   *string `json:"CdnType,omitempty" name:"CdnType"`
	TaskName  *string `json:"TaskName,omitempty" name:"TaskName"`
	Language  *string `json:"Language,omitempty" name:"Language"`
}

func (r *CreateUserUsageDataExportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserUsageDataExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateUserUsageDataExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserUsageDataExportTaskResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	CdnType   *string `json:"CdnType" name:"CdnType"`
	TaskId    *string `json:"TaskId" name:"TaskId"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateUserUsageDataExportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserUsageDataExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUserUsageDataExportTaskRequest struct {
	*ksyunhttp.BaseRequest
	PageSize   *string `json:"PageSize,omitempty" name:"PageSize"`
	PageNumber *string `json:"PageNumber,omitempty" name:"PageNumber"`
}

func (r *GetUserUsageDataExportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserUsageDataExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetUserUsageDataExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetUserUsageDataExportTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId        *string `json:"RequestId" name:"RequestId"`
	UsageDataPerPage struct {
		DataItem []struct {
			TaskId     *string `json:"TaskId"`
			CreateTime *string `json:"CreateTime"`
			UpdateTime *string `json:"UpdateTime"`
			Status     *string `json:"Status"`
			TaskConfig struct {
				StartTime *string `json:"StartTime"`
				EndTime   *string `json:"EndTime"`
			} `json:"TaskConfig"`
			DownloadUrl *string `json:"DownloadUrl"`
		} `json:"DataItem"`
	} `json:"UsageDataPerPage"`
}

func (r *GetUserUsageDataExportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUserUsageDataExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserUsageDataExportTaskRequest struct {
	*ksyunhttp.BaseRequest
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
}

func (r *DeleteUserUsageDataExportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserUsageDataExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteUserUsageDataExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserUsageDataExportTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteUserUsageDataExportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserUsageDataExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDomainUsageDataRequest struct {
	*ksyunhttp.BaseRequest
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" name:"EndTime"`
	Metric    *string `json:"Metric,omitempty" name:"Metric"`
	CdnType   *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains   *string `json:"Domains,omitempty" name:"Domains"`
	Areas     *string `json:"Areas,omitempty" name:"Areas"`
	Interval  *string `json:"Interval,omitempty" name:"Interval"`
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	PeakTime  *string `json:"PeakTime,omitempty" name:"PeakTime"`
	Time      *string `json:"Time,omitempty" name:"Time"`
	Value     *string `json:"Value,omitempty" name:"Value"`
}

func (r *GetDomainUsageDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainUsageDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetDomainUsageDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetDomainUsageDataResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	CdnType   *string `json:"CdnType" name:"CdnType"`
	Domains   *string `json:"Domains" name:"Domains"`
	Areas     *string `json:"Areas" name:"Areas"`
	Metric    []struct {
		Time *string `json:"Time"`
	} `json:"Metric"`
	RequestId *string `json:"RequestId" name:"RequestId"`
	PeakTime  *string `json:"PeakTime" name:"PeakTime"`
}

func (r *GetDomainUsageDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDomainUsageDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUsageDetailDataExportTaskRequest struct {
	*ksyunhttp.BaseRequest
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	EndTime   *string `json:"EndTime,omitempty" name:"EndTime"`
	Type      *string `json:"Type,omitempty" name:"Type"`
	CdnType   *string `json:"CdnType,omitempty" name:"CdnType"`
	Domains   *string `json:"Domains,omitempty" name:"Domains"`
	TaskName  *string `json:"TaskName,omitempty" name:"TaskName"`
	Language  *string `json:"Language,omitempty" name:"Language"`
}

func (r *CreateUsageDetailDataExportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUsageDetailDataExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateUsageDetailDataExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUsageDetailDataExportTaskResponse struct {
	*ksyunhttp.BaseResponse
	StartTime *string `json:"StartTime" name:"StartTime"`
	EndTime   *string `json:"EndTime" name:"EndTime"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateUsageDetailDataExportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUsageDetailDataExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUsageDetailDataExportTaskRequest struct {
	*ksyunhttp.BaseRequest
	PageSize   *string `json:"PageSize,omitempty" name:"PageSize"`
	PageNumber *string `json:"PageNumber,omitempty" name:"PageNumber"`
}

func (r *GetUsageDetailDataExportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUsageDetailDataExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "GetUsageDetailDataExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetUsageDetailDataExportTaskResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *GetUsageDetailDataExportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUsageDetailDataExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsageDetailDataExportTaskRequest struct {
	*ksyunhttp.BaseRequest
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
}

func (r *DeleteUsageDetailDataExportTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUsageDetailDataExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteUsageDetailDataExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsageDetailDataExportTaskResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteUsageDetailDataExportTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUsageDetailDataExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
