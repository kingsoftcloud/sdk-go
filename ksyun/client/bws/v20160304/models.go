package v20160304
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)
type DescribeBandWidthSharesFilter struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Value []*string `json:"Value,omitempty" name:"Value"`
}


type CreateBandWidthShareRequest struct {
    *ksyunhttp.BaseRequest
    LineId *string `json:"LineId,omitempty" name:"LineId"`
    BandWidth *int `json:"BandWidth,omitempty" name:"BandWidth"`
    BandWidthShareName *string `json:"BandWidthShareName,omitempty" name:"BandWidthShareName"`
    ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
    ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
}

func (r *CreateBandWidthShareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBandWidthShareRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateBandWidthShareRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateBandWidthShareResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateBandWidthShareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateBandWidthShareResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBandWidthSharesRequest struct {
    *ksyunhttp.BaseRequest
    ProjectId []*string `json:"ProjectId,omitempty" name:"ProjectId"`
    BandWidthShareId []*string `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
    Filter []*DescribeBandWidthSharesFilter `json:"Filter,omitempty" name:"Filter"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    NextToken *string `json:"NextToken,omitempty" name:"NextToken"`
}

func (r *DescribeBandWidthSharesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBandWidthSharesRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeBandWidthSharesRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBandWidthSharesResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    NextToken *string `json:"NextToken" name:"NextToken"`
	BandWidthShareSet []struct {
		BandWidthShareId *string `json:"BandWidthShareId"`
		BandWidth *int `json:"BandWidth"`
		BandWidthShareName *string `json:"BandWidthShareName"`
		CreateTime *string `json:"CreateTime"`
		LineId *string `json:"LineId"`
		ProjectId *string `json:"ProjectId"`
		AssociateBandWidthShareInfoSet []struct {
					AllocationId *string `json:"AllocationId"`
			} `json:"AssociateBandWidthShareInfoSet"`
		} `json:"BandWidthShareSet"`
}

func (r *DescribeBandWidthSharesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBandWidthSharesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssociateBandWidthShareRequest struct {
    *ksyunhttp.BaseRequest
    BandWidthShareId *string `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
    AllocationId *string `json:"AllocationId,omitempty" name:"AllocationId"`
}

func (r *AssociateBandWidthShareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateBandWidthShareRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AssociateBandWidthShareRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AssociateBandWidthShareResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
}

func (r *AssociateBandWidthShareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssociateBandWidthShareResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateBandWidthShareRequest struct {
    *ksyunhttp.BaseRequest
    BandWidthShareId *string `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
    AllocationId *string `json:"AllocationId,omitempty" name:"AllocationId"`
    BandWidth *int `json:"BandWidth,omitempty" name:"BandWidth"`
}

func (r *DisassociateBandWidthShareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateBandWidthShareRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DisassociateBandWidthShareRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DisassociateBandWidthShareResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
}

func (r *DisassociateBandWidthShareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisassociateBandWidthShareResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBandWidthShareRequest struct {
    *ksyunhttp.BaseRequest
    BandWidthShareId *string `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
    BandWidth *int `json:"BandWidth,omitempty" name:"BandWidth"`
    BandWidthShareName *string `json:"BandWidthShareName,omitempty" name:"BandWidthShareName"`
}

func (r *ModifyBandWidthShareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBandWidthShareRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyBandWidthShareRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ModifyBandWidthShareResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyBandWidthShareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyBandWidthShareResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBandWidthShareRequest struct {
    *ksyunhttp.BaseRequest
    BandWidthShareId *string `json:"BandWidthShareId,omitempty" name:"BandWidthShareId"`
}

func (r *DeleteBandWidthShareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBandWidthShareRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteBandWidthShareRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteBandWidthShareResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
    Return *bool `json:"Return" name:"Return"`
}

func (r *DeleteBandWidthShareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteBandWidthShareResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

