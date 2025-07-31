package v20160304

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type DescribeCfwAvRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *DescribeCfwAvRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeCfwAvResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	CfwAvs    []struct {
		AvId        *string `json:"AvId" name:"AvId"`
		Protocol    *string `json:"Protocol" name:"Protocol"`
		ProtectType *string `json:"ProtectType" name:"ProtectType"`
		Status      *string `json:"Status" name:"Status"`
		CreateTime  *string `json:"CreateTime" name:"CreateTime"`
	} `json:"CfwAvs"`
}

func (r *DescribeCfwAvResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCfwAvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudFireWallInstanceRequest struct {
	*ksyunhttp.BaseRequest
	CfwInstanceIds []*string `json:"CfwInstanceIds,omitempty" name:"CfwInstanceIds"`
}

func (r *DescribeCloudFireWallInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeCloudFireWallInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                 *string `json:"requestId" name:"requestId"`
	CloudFireWallInstanceList []struct {
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		InstanceType *string `json:"InstanceType" name:"InstanceType"`
		Bandwidth    *int    `json:"Bandwidth" name:"Bandwidth"`
		Status       *int    `json:"Status" name:"Status"`
		TotalEipNum  *int    `json:"TotalEipNum" name:"TotalEipNum"`
		UsedEipNum   *int    `json:"UsedEipNum" name:"UsedEipNum"`
		IpsStatus    *int    `json:"IpsStatus" name:"IpsStatus"`
		AvStatus     *int    `json:"AvStatus" name:"AvStatus"`
		CreateTime   *string `json:"CreateTime" name:"CreateTime"`
	} `json:"CloudFireWallInstanceList"`
}

func (r *DescribeCloudFireWallInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCloudFireWallInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
