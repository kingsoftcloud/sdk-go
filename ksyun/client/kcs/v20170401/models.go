package v20170401
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type DescribeCacheReadonlyNodeRequest struct {
    *ksyunhttp.BaseRequest
    CacheId *string `json:"CacheId,omitempty" name:"CacheId"`
    AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
}

func (r *DescribeCacheReadonlyNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCacheReadonlyNodeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCacheReadonlyNodeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheReadonlyNodeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	Data []struct {
		InstanceId *string `json:"InstanceId"`
		Name *string `json:"Name"`
		Ip *string `json:"Ip"`
		Status *string `json:"Status"`
		CreateTime *string `json:"CreateTime"`
	} `json:"Data"`
}

func (r *DescribeCacheReadonlyNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCacheReadonlyNodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddCacheSlaveNodeRequest struct {
    *ksyunhttp.BaseRequest
    AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
    CacheId *string `json:"CacheId,omitempty" name:"CacheId"`
    SlaveVip *string `json:"SlaveVip,omitempty" name:"SlaveVip"`
}

func (r *AddCacheSlaveNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddCacheSlaveNodeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "AddCacheSlaveNodeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type AddCacheSlaveNodeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
	Data struct {
		NodeId *string `json:"NodeId"`
	} `json:"Data"`
}

func (r *AddCacheSlaveNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddCacheSlaveNodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCacheSlaveNodeRequest struct {
    *ksyunhttp.BaseRequest
    AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
    CacheId *string `json:"CacheId,omitempty" name:"CacheId"`
    NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *DeleteCacheSlaveNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCacheSlaveNodeRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteCacheSlaveNodeRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteCacheSlaveNodeResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteCacheSlaveNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteCacheSlaveNodeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

