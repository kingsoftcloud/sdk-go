package v20170401
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type DeleteCacheSlaveNodeRequest struct {
	*ksyunhttp.BaseRequest
	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	CacheId       *string `json:"CacheId,omitempty" name:"CacheId"`
	NodeId        *string `json:"NodeId,omitempty" name:"NodeId"`
}

func (r *DeleteCacheSlaveNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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

