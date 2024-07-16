package v20201231
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)
type CreateClusterManagedClusterMultiMaster struct {
    SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
    SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
}

type CreateClusterInstanceForNode struct {
    NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`
NodeConfig   []struct {
                Para *string `json:"Para,omitempty" name:"Para"`
                    } `json:"NodeConfig,omitempty" name:"NodeConfig"`
}

type CreateClusterNodeConfig struct {
    Para *string `json:"Para,omitempty" name:"Para"`
}

type CreateClusterLabel struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateClusterKubelet struct {
    CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
}

type CreateClusterTaints struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
    Effect *string `json:"Effect,omitempty" name:"Effect"`
}

type CreateClusterExistedInstanceForEpc struct {
    NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`
EpcConfig   []struct {
                Para *string `json:"Para,omitempty" name:"Para"`
                    } `json:"EpcConfig,omitempty" name:"EpcConfig"`
}

type CreateClusterEpcConfig struct {
    Para *string `json:"Para,omitempty" name:"Para"`
}

type CreateClusterLabel struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
}

type CreateClusterKubelet struct {
    CustomArg *string `json:"CustomArg,omitempty" name:"CustomArg"`
}

type CreateClusterTaint struct {
    Key *string `json:"Key,omitempty" name:"Key"`
    Value *string `json:"Value,omitempty" name:"Value"`
    Effect *string `json:"Effect,omitempty" name:"Effect"`
}

type CreateClusterComponent struct {
    Name *string `json:"Name,omitempty" name:"Name"`
    Config *string `json:"Config,omitempty" name:"Config"`
}


type CreateClusterRequest struct {
    *ksyunhttp.BaseRequest
    ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
    ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
    ClusterManageMode *string `json:"ClusterManageMode,omitempty" name:"ClusterManageMode"`
    ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`
    VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
    PodCidr *string `json:"PodCidr,omitempty" name:"PodCidr"`
    ServiceCidr *string `json:"ServiceCidr,omitempty" name:"ServiceCidr"`
    NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
    K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`
    ReserveSubnetId *string `json:"ReserveSubnetId,omitempty" name:"ReserveSubnetId"`
    PublicApiServer *string `json:"PublicApiServer,omitempty" name:"PublicApiServer"`
    ExposePublicApiServer *bool `json:"ExposePublicApiServer,omitempty" name:"ExposePublicApiServer"`
    MaxPodPerNode *string `json:"MaxPodPerNode,omitempty" name:"MaxPodPerNode"`
    MasterEtcdSeparate *bool `json:"MasterEtcdSeparate,omitempty" name:"MasterEtcdSeparate"`
    ManagedClusterMultiMaster []*CreateClusterManagedClusterMultiMaster `json:"ManagedClusterMultiMaster,omitempty" name:"ManagedClusterMultiMaster"`
    InstanceForNode []*CreateClusterInstanceForNode `json:"InstanceForNode,omitempty" name:"InstanceForNode"`
    ExistedInstanceForEpc []*CreateClusterExistedInstanceForEpc `json:"ExistedInstanceForEpc,omitempty" name:"ExistedInstanceForEpc"`
    Component []*CreateClusterComponent `json:"Component,omitempty" name:"Component"`
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
    ClusterId *string `json:"ClusterId" name:"ClusterId"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClusterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

