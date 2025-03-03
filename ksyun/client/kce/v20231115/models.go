package v20231115

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type DescribeClusterFilter struct {
	Name  *string   `json:"Name,omitempty" name:"Name"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeClusterRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId  *string                  `json:"ClusterId,omitempty" name:"ClusterId"`
	Marker     *int                     `json:"Marker,omitempty" name:"Marker"`
	MaxResults *int                     `json:"MaxResults,omitempty" name:"MaxResults"`
	Search     *string                  `json:"Search,omitempty" name:"Search"`
	Filter     []*DescribeClusterFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	MaxResults *int    `json:"MaxResults" name:"MaxResults"`
	Marker     *int    `json:"Marker" name:"Marker"`
	ClusterSet []struct {
		ClusterId          *string `json:"ClusterId" name:"ClusterId"`
		ClusterName        *string `json:"ClusterName" name:"ClusterName"`
		ClusterType        *string `json:"ClusterType" name:"ClusterType"`
		ClusterManageMode  *string `json:"ClusterManageMode" name:"ClusterManageMode"`
		K8sVersion         *string `json:"K8sVersion" name:"K8sVersion"`
		ClusterDesc        *string `json:"ClusterDesc" name:"ClusterDesc"`
		PodCidr            *string `json:"PodCidr" name:"PodCidr"`
		ServiceCidr        *string `json:"ServiceCidr" name:"ServiceCidr"`
		VpcId              *string `json:"VpcId" name:"VpcId"`
		VpcCidr            *string `json:"VpcCidr" name:"VpcCidr"`
		Status             *string `json:"Status" name:"Status"`
		NodeNum            *int    `json:"NodeNum" name:"NodeNum"`
		NormalNodeNum      *int    `json:"NormalNodeNum" name:"NormalNodeNum"`
		NetworkType        *string `json:"NetworkType" name:"NetworkType"`
		MaxPodPerNode      *int    `json:"MaxPodPerNode" name:"MaxPodPerNode"`
		MasterEtcdSeparate *bool   `json:"MasterEtcdSeparate" name:"MasterEtcdSeparate"`
		OrderType          *int    `json:"OrderType" name:"OrderType"`
		ServiceEndTime     *string `json:"ServiceEndTime" name:"ServiceEndTime"`
		EnableKMSE         *bool   `json:"EnableKMSE" name:"EnableKMSE"`
		CreateTime         *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime         *string `json:"UpdateTime" name:"UpdateTime"`
	} `json:"ClusterSet"`
}

func (r *DescribeClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterDelProtectionRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId           *string `json:"ClusterId,omitempty" name:"ClusterId"`
	EnableDelProtection *bool   `json:"EnableDelProtection,omitempty" name:"EnableDelProtection"`
}

func (r *UpdateClusterDelProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateClusterDelProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateClusterDelProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterDelProtectionResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateClusterDelProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateClusterDelProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
