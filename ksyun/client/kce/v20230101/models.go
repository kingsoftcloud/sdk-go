package v20230101

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type DescribeComponentParamsComponents struct {
	Type         *string `json:"Type,omitempty" name:"Type"`
	ParamVersion *string `json:"ParamVersion,omitempty" name:"ParamVersion"`
}

type DescribeEventLogsRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	NodeId      *string `json:"NodeId,omitempty" name:"NodeId"`
	NodeName    *string `json:"NodeName,omitempty" name:"NodeName"`
	Inner       *bool   `json:"Inner,omitempty" name:"Inner"`
	Marker      *int    `json:"Marker,omitempty" name:"Marker"`
	MaxResults  *int    `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeEventLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeEventLogsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		ClusterId  *string `json:"ClusterId" name:"ClusterId"`
		MaxResults *int    `json:"MaxResults" name:"MaxResults"`
		TotalCount *int    `json:"TotalCount" name:"TotalCount"`
		EventLogs  struct {
			RuntimeInfo struct {
				RuntimeName *string `json:"RuntimeName" name:"RuntimeName"`
				RuntimeIP   *string `json:"RuntimeIP" name:"RuntimeIP"`
				NodeIP      *string `json:"NodeIP" name:"NodeIP"`
			} `json:"RuntimeInfo"`
			UserInfo struct {
				AccountId *string `json:"AccountId" name:"AccountId"`
				Region    *string `json:"Region" name:"Region"`
			} `json:"UserInfo"`
			EventInfo struct {
				EventId     *string `json:"EventId" name:"EventId"`
				ClusterId   *string `json:"ClusterId" name:"ClusterId"`
				EventType   *string `json:"EventType" name:"EventType"`
				Level       *string `json:"Level" name:"Level"`
				CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
				Content     *string `json:"Content" name:"Content"`
				Category    *int    `json:"Category" name:"Category"`
			} `json:"EventInfo"`
		} `json:"EventLogs" name:"EventLogs"`
	} `json:"Data"`
}

func (r *DescribeEventLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAddonInstanceRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *CreateAddonInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateAddonInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		ClusterId   *string   `json:"ClusterId" name:"ClusterId"`
		InstanceIds []*string `json:"InstanceIds" name:"InstanceIds"`
	} `json:"Data"`
}

func (r *CreateAddonInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAddonInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAddonInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	AddonId     *string `json:"AddonId,omitempty" name:"AddonId"`
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteAddonInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteAddonInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequesetId *string `json:"RequesetId" name:"RequesetId"`
	Data       struct {
	} `json:"Data"`
	ClusterId *string `json:"ClusterId" name:"ClusterId"`
	AddonId   *string `json:"AddonId" name:"AddonId"`
}

func (r *DeleteAddonInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAddonInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddonInstancesRequest struct {
	*ksyunhttp.BaseRequest
	CulsterId   *string   `json:"CulsterId,omitempty" name:"CulsterId"`
	ClusterName *string   `json:"ClusterName,omitempty" name:"ClusterName"`
	Name        *string   `json:"Name,omitempty" name:"Name"`
	AddonIds    []*string `json:"AddonIds,omitempty" name:"AddonIds"`
}

func (r *DescribeAddonInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAddonInstancesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Addons []struct {
			ClusterId   *string `json:"ClusterId" name:"ClusterId"`
			AddonId     *string `json:"AddonId" name:"AddonId"`
			InstanceId  *string `json:"InstanceId" name:"InstanceId"`
			Type        *string `json:"Type" name:"Type"`
			ToDelete    *bool   `json:"ToDelete" name:"ToDelete"`
			Phase       *string `json:"Phase" name:"Phase"`
			CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
			UpdatedTime *string `json:"UpdatedTime" name:"UpdatedTime"`
		} `json:"Addons" name:"Addons"`
	} `json:"Data"`
	Args struct {
	} `json:"Args"`
}

func (r *DescribeAddonInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddonInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAddonListRequest struct {
	*ksyunhttp.BaseRequest
	Name       *string `json:"Name,omitempty" name:"Name"`
	MaxResults *int    `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker     *int    `json:"Marker,omitempty" name:"Marker"`
}

func (r *DescribeAddonListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeAddonListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		AddonList []struct {
			CompatibleK8SVersion struct {
				Max            *string `json:"Max" name:"Max"`
				Min            *string `json:"Min" name:"Min"`
				DescriptionCn  *string `json:"DescriptionCn" name:"DescriptionCn"`
				Description    *string `json:"Description" name:"Description"`
				Category       *string `json:"Category" name:"Category"`
				SubCategory    *string `json:"SubCategory" name:"SubCategory"`
				AddonId        *string `json:"AddonId" name:"AddonId"`
				AddonVersion   *string `json:"AddonVersion" name:"AddonVersion"`
				Name           *string `json:"Name" name:"Name"`
				DefaultInstall *bool   `json:"DefaultInstall" name:"DefaultInstall"`
			} `json:"CompatibleK8SVersion"`
		} `json:"AddonList" name:"AddonList"`
	} `json:"Data"`
}

func (r *DescribeAddonListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAddonListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentParamsRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId  *string                              `json:"ClusterId,omitempty" name:"ClusterId"`
	Components []*DescribeComponentParamsComponents `json:"Components,omitempty" name:"Components"`
	Marker     *int                                 `json:"Marker,omitempty" name:"Marker"`
	MaxResults *int                                 `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeComponentParamsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeComponentParamsResponse struct {
	*ksyunhttp.BaseResponse
	RequeestId *string `json:"RequeestId" name:"RequeestId"`
	Data       struct {
		ClusterId *string `json:"ClusterId" name:"ClusterId"`
	} `json:"Data"`
	Components []struct {
		Type    *string `json:"Type" name:"Type"`
		Version *string `json:"Version" name:"Version"`
		Args    *string `json:"Args" name:"Args"`
	} `json:"Components"`
}

func (r *DescribeComponentParamsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentParamsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNetworkRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *DescribeNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeNetworkResponse struct {
	*ksyunhttp.BaseResponse
	ClusterId *string `json:"ClusterId" name:"ClusterId"`
	NetworkId *string `json:"NetworkId" name:"NetworkId"`
	PublicSLB struct {
		Phase  *string `json:"Phase" name:"Phase"`
		Reason *string `json:"Reason" name:"Reason"`
	} `json:"PublicSLB"`
	PrivateSLB struct {
		SLBId  *string `json:"SLBId" name:"SLBId"`
		SLBIp  *string `json:"SLBIp" name:"SLBIp"`
		Phase  *string `json:"Phase" name:"Phase"`
		Reason *string `json:"Reason" name:"Reason"`
	} `json:"PrivateSLB"`
	PrivateLink struct {
		LinkIp   *string `json:"LinkIp" name:"LinkIp"`
		LinkPort *string `json:"LinkPort" name:"LinkPort"`
		Phase    *string `json:"Phase" name:"Phase"`
		Reason   *string `json:"Reason" name:"Reason"`
	} `json:"PrivateLink"`
	EIP struct {
		EIPId  *string `json:"EIPId" name:"EIPId"`
		Phase  *string `json:"Phase" name:"Phase"`
		Reason *string `json:"Reason" name:"Reason"`
	} `json:"EIP"`
	PublicAccess *bool   `json:"PublicAccess" name:"PublicAccess"`
	CreatedTime  *string `json:"CreatedTime" name:"CreatedTime"`
	UpdatedTime  *string `json:"UpdatedTime" name:"UpdatedTime"`
}

func (r *DescribeNetworkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeComponentsRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	ClusterName *string   `json:"ClusterName,omitempty" name:"ClusterName"`
	NodeNames   []*string `json:"NodeNames,omitempty" name:"NodeNames"`
	NodeIds     *string   `json:"NodeIds,omitempty" name:"NodeIds"`
	Marker      *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults  *int      `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeNodeComponentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeNodeComponentsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		ClusterId      *string `json:"ClusterId" name:"ClusterId"`
		NodeComponents []struct {
			NodeId          *string `json:"NodeId" name:"NodeId"`
			ComponentStatus []struct {
				Type        *string `json:"Type" name:"Type"`
				CurVersion  *string `json:"CurVersion" name:"CurVersion"`
				SpecVersion *string `json:"SpecVersion" name:"SpecVersion"`
			} `json:"ComponentStatus"`
		} `json:"NodeComponents" name:"NodeComponents"`
	} `json:"Data"`
}

func (r *DescribeNodeComponentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeComponentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeComponentListRequest struct {
	*ksyunhttp.BaseRequest
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`
}

func (r *DescribeComponentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeComponentListResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeComponentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeComponentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
