package v20240701

import (
	"encoding/json"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type DescribeCacheByRoleRequest struct {
	*ksyunhttp.BaseRequest
	CacheId *string `json:"CacheId,omitempty" name:"CacheId"`
	Role    *string `json:"Role,omitempty" name:"Role"`
}

func (r *DescribeCacheByRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheByRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DescribeCacheByRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCacheByRoleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		MasterNodes []struct {
			InstanceId   *string `json:"InstanceId" name:"InstanceId"`
			Name         *string `json:"Name" name:"Name"`
			Port         *int    `json:"Port" name:"Port"`
			Ip           *string `json:"Ip" name:"Ip"`
			Status       *string `json:"Status" name:"Status"`
			CreateTime   *string `json:"CreateTime" name:"CreateTime"`
			InstanceRole *string `json:"InstanceRole" name:"InstanceRole"`
			Az           *string `json:"Az" name:"Az"`
			ShardId      *string `json:"ShardId" name:"ShardId"`
		} `json:"MasterNodes" name:"MasterNodes"`
		SlavesNodes []struct {
			InstanceId   *string `json:"InstanceId" name:"InstanceId"`
			Name         *string `json:"Name" name:"Name"`
			Port         *int    `json:"Port" name:"Port"`
			Ip           *string `json:"Ip" name:"Ip"`
			Status       *string `json:"Status" name:"Status"`
			CreateTime   *string `json:"CreateTime" name:"CreateTime"`
			InstanceRole *string `json:"InstanceRole" name:"InstanceRole"`
			Az           *string `json:"Az" name:"Az"`
			ShardId      *string `json:"ShardId" name:"ShardId"`
		} `json:"SlavesNodes" name:"SlavesNodes"`
		ProxyNodes []struct {
			InstanceId   *string `json:"InstanceId" name:"InstanceId"`
			Name         *string `json:"Name" name:"Name"`
			Port         *int    `json:"Port" name:"Port"`
			Ip           *string `json:"Ip" name:"Ip"`
			Status       *string `json:"Status" name:"Status"`
			CreateTime   *string `json:"CreateTime" name:"CreateTime"`
			InstanceRole *string `json:"InstanceRole" name:"InstanceRole"`
			Az           *string `json:"Az" name:"Az"`
			ShardId      *string `json:"ShardId" name:"ShardId"`
		} `json:"ProxyNodes" name:"ProxyNodes"`
	} `json:"Data"`
}

func (r *DescribeCacheByRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCacheByRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
