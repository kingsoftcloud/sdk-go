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
		ClusterId          *string `json:"ClusterId"`
		ClusterName        *string `json:"ClusterName"`
		ClusterType        *string `json:"ClusterType"`
		ClusterManageMode  *string `json:"ClusterManageMode"`
		K8sVersion         *string `json:"K8sVersion"`
		ClusterDesc        *string `json:"ClusterDesc"`
		PodCidr            *string `json:"PodCidr"`
		ServiceCidr        *string `json:"ServiceCidr"`
		VpcId              *string `json:"VpcId"`
		VpcCidr            *string `json:"VpcCidr"`
		Status             *string `json:"Status"`
		NodeNum            *int    `json:"NodeNum"`
		NormalNodeNum      *int    `json:"NormalNodeNum"`
		NetworkType        *string `json:"NetworkType"`
		MaxPodPerNode      *int    `json:"MaxPodPerNode"`
		MasterEtcdSeparate *bool   `json:"MasterEtcdSeparate"`
		OrderType          *int    `json:"OrderType"`
		ServiceEndTime     *string `json:"ServiceEndTime"`
		EnableKMSE         *bool   `json:"EnableKMSE"`
		CreateTime         *string `json:"CreateTime"`
		UpdateTime         *string `json:"UpdateTime"`
	} `json:"ClusterSet"`
}

func (r *DescribeClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
