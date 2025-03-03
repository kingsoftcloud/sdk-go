package v20230101

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type CreateClusterComponents struct {
	Type *string `json:"Type,omitempty" name:"Type"`
	Args *string `json:"Args,omitempty" name:"Args"`
}

type CreateClusterPrivateRegistries struct {
	Url      *string `json:"Url,omitempty" name:"Url"`
	User     *string `json:"User,omitempty" name:"User"`
	Password *string `json:"Password,omitempty" name:"Password"`
}

type CreateClusterNodeInstanceSet struct {
	Components []struct {
		Type *string `json:"Type,omitempty" name:"Type"`
		Args *string `json:"Args,omitempty" name:"Args"`
	} `json:"Components,omitempty" name:"Components"`
}

type CreateClusterLabels struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateClusterTaints struct {
	Key    *string `json:"Key,omitempty" name:"Key"`
	Value  *string `json:"Value,omitempty" name:"Value"`
	Effect *string `json:"Effect,omitempty" name:"Effect"`
}

type CreateClusterAddons struct {
	Name         *string `json:"Name,omitempty" name:"Name"`
	Type         *string `json:"Type,omitempty" name:"Type"`
	AddonVersion *string `json:"AddonVersion,omitempty" name:"AddonVersion"`
}

type CreateClusterBastionHosts struct {
	Address  *string `json:"Address,omitempty" name:"Address"`
	Port     *string `json:"Port,omitempty" name:"Port"`
	UserName *string `json:"UserName,omitempty" name:"UserName"`
	Password *string `json:"Password,omitempty" name:"Password"`
	Order    *int    `json:"Order,omitempty" name:"Order"`
	Script   *string `json:"Script,omitempty" name:"Script"`
}

type AddNodesNodeInstanceSet struct {
	Components []struct {
		Type *string `json:"Type,omitempty" name:"Type"`
		Args *string `json:"Args,omitempty" name:"Args"`
	} `json:"Components,omitempty" name:"Components"`
}

type AddNodesInstanceLabel struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}

type AddNodesLabels struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}

type AddNodesTaints struct {
	Key    *string `json:"Key,omitempty" name:"Key"`
	Value  *string `json:"Value,omitempty" name:"Value"`
	Effect *string `json:"Effect,omitempty" name:"Effect"`
}

type AddNodesComponents struct {
	Type *string `json:"Type,omitempty" name:"Type"`
	Args *string `json:"Args,omitempty" name:"Args"`
}

type DescribeComponentParamsComponents struct {
	Type         *string `json:"Type,omitempty" name:"Type"`
	ParamVersion *string `json:"ParamVersion,omitempty" name:"ParamVersion"`
}

type CreateAddonInstanceAddons struct {
	Name         *string `json:"Name,omitempty" name:"Name"`
	Type         *string `json:"Type,omitempty" name:"Type"`
	Config       *string `json:"Config,omitempty" name:"Config"`
	AddonVersion *string `json:"AddonVersion,omitempty" name:"AddonVersion"`
}

type CreateClusterRequest struct {
	*ksyunhttp.BaseRequest
	ClusterName       *string                         `json:"ClusterName,omitempty" name:"ClusterName"`
	ClusterDesc       *string                         `json:"ClusterDesc,omitempty" name:"ClusterDesc"`
	ClusterManageMode *string                         `json:"ClusterManageMode,omitempty" name:"ClusterManageMode"`
	ProjectId         *string                         `json:"ProjectId,omitempty" name:"ProjectId"`
	NodeInstanceSet   []*CreateClusterNodeInstanceSet `json:"NodeInstanceSet,omitempty" name:"NodeInstanceSet"`
	Addons            []*CreateClusterAddons          `json:"Addons,omitempty" name:"Addons"`
}

func (r *CreateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
	} `json:"Data"`
	ClusterId *string `json:"ClusterId" name:"ClusterId"`
}

func (r *CreateClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersRequest struct {
	*ksyunhttp.BaseRequest
	ClusterNames []*string `json:"ClusterNames,omitempty" name:"ClusterNames"`
	ClusterIds   []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	MaxResults   *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker       *int      `json:"Marker,omitempty" name:"Marker"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		MaxResults *int `json:"MaxResults" name:"MaxResults"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
	} `json:"Data"`
	ClusterSet []struct {
	} `json:"ClusterSet"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRequest struct {
	*ksyunhttp.BaseRequest
	ClusterName    *string `json:"ClusterName,omitempty" name:"ClusterName"`
	ClusterId      *string `json:"ClusterId,omitempty" name:"ClusterId"`
	InstanceDelete *bool   `json:"InstanceDelete,omitempty" name:"InstanceDelete"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterResponse struct {
	*ksyunhttp.BaseResponse
	Request *string `json:"Request" name:"Request"`
	Data    struct {
		ClusterId *string `json:"ClusterId" name:"ClusterId"`
	} `json:"Data"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	ClusterName *string   `json:"ClusterName,omitempty" name:"ClusterName"`
	ClusterDesc *string   `json:"ClusterDesc,omitempty" name:"ClusterDesc"`
	SANs        []*string `json:"SANs,omitempty" name:"SANs"`
}

func (r *ModifyClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		ClusterId *string `json:"ClusterId" name:"ClusterId"`
	} `json:"Data"`
}

func (r *ModifyClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadClusterConfigRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	Type        *string `json:"Type,omitempty" name:"Type"`
}

func (r *DownloadClusterConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadClusterConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DownloadClusterConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DownloadClusterConfigResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		ClusterId *string `json:"ClusterId" name:"ClusterId"`
		Type      *string `json:"Type" name:"Type"`
		Config    *string `json:"Config" name:"Config"`
	} `json:"Data"`
}

func (r *DownloadClusterConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DownloadClusterConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddNodesRequest struct {
	*ksyunhttp.BaseRequest
	ClusterName     *string                    `json:"ClusterName,omitempty" name:"ClusterName"`
	ClusterId       *string                    `json:"ClusterId,omitempty" name:"ClusterId"`
	NodeInstanceSet []*AddNodesNodeInstanceSet `json:"NodeInstanceSet,omitempty" name:"NodeInstanceSet"`
}

func (r *AddNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddNodesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		InstanceSet []struct {
			MachineId *string `json:"MachineId" name:"MachineId"`
			Code      *int    `json:"Code" name:"Code"`
		} `json:"InstanceSet" name:"InstanceSet"`
	} `json:"Data"`
}

func (r *AddNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodesRequest struct {
	*ksyunhttp.BaseRequest
	ClusterName *string   `json:"ClusterName,omitempty" name:"ClusterName"`
	ClusterId   *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	NodeNames   []*string `json:"NodeNames,omitempty" name:"NodeNames"`
	NodeIds     []*string `json:"NodeIds,omitempty" name:"NodeIds"`
	Marker      *int      `json:"Marker,omitempty" name:"Marker"`
	MaxResults  *int      `json:"MaxResults,omitempty" name:"MaxResults"`
	Search      *string   `json:"Search,omitempty" name:"Search"`
}

func (r *DescribeNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		MaxResults *int `json:"MaxResults" name:"MaxResults"`
		Marker     *int `json:"Marker" name:"Marker"`
		TotalCount *int `json:"TotalCount" name:"TotalCount"`
	} `json:"Data"`
	InstanceSet struct {
	} `json:"InstanceSet"`
}

func (r *DescribeNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNodeRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId      *string   `json:"ClusterId,omitempty" name:"ClusterId"`
	ClusterName    *string   `json:"ClusterName,omitempty" name:"ClusterName"`
	NodeIds        []*string `json:"NodeIds,omitempty" name:"NodeIds"`
	NodeNames      *string   `json:"NodeNames,omitempty" name:"NodeNames"`
	InstanceDelete *bool     `json:"InstanceDelete,omitempty" name:"InstanceDelete"`
}

func (r *DeleteNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNodeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		ClusterId   *string `json:"ClusterId" name:"ClusterId"`
		DeleteCount *int    `json:"DeleteCount" name:"DeleteCount"`
	} `json:"Data"`
}

func (r *DeleteNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNodeRequest struct {
	*ksyunhttp.BaseRequest
	ClusteId    *string   `json:"ClusteId,omitempty" name:"ClusteId"`
	ClusterName *string   `json:"ClusterName,omitempty" name:"ClusterName"`
	NodeId      *string   `json:"NodeId,omitempty" name:"NodeId"`
	NodeName    *string   `json:"NodeName,omitempty" name:"NodeName"`
	Components  []*string `json:"Components,omitempty" name:"Components"`
}

func (r *ModifyNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ModifyNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNodeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNodeResponse) FromJsonString(s string) error {
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

func (r *DescribeComponentListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeComponentListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeNodeComponentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeNodeComponentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type DescribeNetworkRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId   *string `json:"ClusterId,omitempty" name:"ClusterId"`
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *DescribeNetworkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeComponentParamsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeComponentParamsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeAddonListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAddonListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeAddonInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeAddonInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

func (r *DeleteAddonInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteAddonInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
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

type CreateAddonInstanceRequest struct {
	*ksyunhttp.BaseRequest
	ClusterId *string                      `json:"ClusterId,omitempty" name:"ClusterId"`
	Addons    []*CreateAddonInstanceAddons `json:"Addons,omitempty" name:"Addons"`
}

func (r *CreateAddonInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAddonInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateAddonInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAddonInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		ClusterId   *string `json:"ClusterId" name:"ClusterId"`
		InstanceIds []struct {
		} `json:"InstanceIds" name:"InstanceIds"`
	} `json:"Data"`
}

func (r *CreateAddonInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAddonInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

func (r *DescribeEventLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeEventLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventLogsResponse struct {
	*ksyunhttp.BaseResponse
}

func (r *DescribeEventLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterVersionListRequest struct {
	*ksyunhttp.BaseRequest
	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`
}

func (r *DescribeClusterVersionListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterVersionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeClusterVersionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterVersionListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
	} `json:"Data"`
}

func (r *DescribeClusterVersionListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterVersionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
