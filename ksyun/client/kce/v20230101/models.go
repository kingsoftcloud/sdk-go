package v20230101
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)
type DescribeComponentParamsComponents struct {
    Type *string `json:"Type,omitempty" name:"Type"`
    ParamVersion *string `json:"ParamVersion,omitempty" name:"ParamVersion"`
}


type DescribeEventLogsRequest struct {
    *ksyunhttp.BaseRequest
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
    NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
    NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
    Inner *bool `json:"Inner,omitempty" name:"Inner"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
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
    RequestId *string `json:"RequestId" name:"RequestId"`
	Data struct {
		ClusterId *string `json:"ClusterId"`
		MaxResults *int `json:"MaxResults"`
		TotalCount *int `json:"TotalCount"`
		EventLogs struct {
			RuntimeInfo struct {
				RuntimeName *string `json:"RuntimeName"`
				RuntimeIP *string `json:"RuntimeIP"`
				NodeIP *string `json:"NodeIP"`
			} `json:"RuntimeInfo"`
			UserInfo struct {
				AccountId *string `json:"AccountId"`
				Region *string `json:"Region"`
			} `json:"UserInfo"`
			EventInfo struct {
				EventId *string `json:"EventId"`
				ClusterId *string `json:"ClusterId"`
				EventType *string `json:"EventType"`
				Level *string `json:"Level"`
				CreatedTime *string `json:"CreatedTime"`
				Content *string `json:"Content"`
				Category *int `json:"Category"`
			} `json:"EventInfo"`
		} `json:"EventLogs"`
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
	Data struct {
		ClusterId *string `json:"ClusterId"`
		InstanceIds []struct {
			} `json:"InstanceIds"`
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
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    AddonId *string `json:"AddonId,omitempty" name:"AddonId"`
    InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
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
	Data struct {
	} `json:"Data"`
    ClusterId *string `json:"ClusterId" name:"ClusterId"`
    AddonId *string `json:"AddonId" name:"AddonId"`
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
    CulsterId *string `json:"CulsterId,omitempty" name:"CulsterId"`
    ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
    Name *string `json:"Name,omitempty" name:"Name"`
    AddonIds []*string `json:"AddonIds,omitempty" name:"AddonIds"`
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
	Data struct {
		Addons []struct {
					ClusterId *string `json:"ClusterId"`
					AddonId *string `json:"AddonId"`
					InstanceId *string `json:"InstanceId"`
					Type *string `json:"Type"`
					ToDelete *bool `json:"ToDelete"`
					Phase *string `json:"Phase"`
					CreatedTime *string `json:"CreatedTime"`
					UpdatedTime *string `json:"UpdatedTime"`
			} `json:"Addons"`
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
    Name *string `json:"Name,omitempty" name:"Name"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
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
	Data struct {
		AddonList []struct {
				CompatibleK8SVersion struct {
					Max *string `json:"Max"`
					Min *string `json:"Min"`
					DescriptionCn *string `json:"DescriptionCn"`
					Description *string `json:"Description"`
					Category *string `json:"Category"`
					SubCategory *string `json:"SubCategory"`
					AddonId *string `json:"AddonId"`
					AddonVersion *string `json:"AddonVersion"`
					Name *string `json:"Name"`
					DefaultInstall *bool `json:"DefaultInstall"`
				} `json:"CompatibleK8SVersion"`
			} `json:"AddonList"`
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
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    Components []*DescribeComponentParamsComponents `json:"Components,omitempty" name:"Components"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
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
	Data struct {
		ClusterId *string `json:"ClusterId"`
	} `json:"Data"`
	Components []struct {
		Type *string `json:"Type"`
		Version *string `json:"Version"`
		Args *string `json:"Args"`
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
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
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
		Phase *string `json:"Phase"`
		Reason *string `json:"Reason"`
	} `json:"PublicSLB"`
	PrivateSLB struct {
		SLBId *string `json:"SLBId"`
		SLBIp *string `json:"SLBIp"`
		Phase *string `json:"Phase"`
		Reason *string `json:"Reason"`
	} `json:"PrivateSLB"`
	PrivateLink struct {
		LinkIp *string `json:"LinkIp"`
		LinkPort *string `json:"LinkPort"`
		Phase *string `json:"Phase"`
		Reason *string `json:"Reason"`
	} `json:"PrivateLink"`
	EIP struct {
		EIPId *string `json:"EIPId"`
		Phase *string `json:"Phase"`
		Reason *string `json:"Reason"`
	} `json:"EIP"`
    PublicAccess *bool `json:"PublicAccess" name:"PublicAccess"`
    CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
    UpdatedTime *string `json:"UpdatedTime" name:"UpdatedTime"`
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
    ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
    ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
    NodeNames []*string `json:"NodeNames,omitempty" name:"NodeNames"`
    NodeIds *string `json:"NodeIds,omitempty" name:"NodeIds"`
    Marker *int `json:"Marker,omitempty" name:"Marker"`
    MaxResults *int `json:"MaxResults,omitempty" name:"MaxResults"`
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
	Data struct {
		ClusterId *string `json:"ClusterId"`
		NodeComponents []struct {
					NodeId *string `json:"NodeId"`
				ComponentStatus []struct {
					Type *string `json:"Type"`
					CurVersion *string `json:"CurVersion"`
					SpecVersion *string `json:"SpecVersion"`
				} `json:"ComponentStatus"`
			} `json:"NodeComponents"`
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

