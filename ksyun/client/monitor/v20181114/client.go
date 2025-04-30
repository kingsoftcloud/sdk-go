package v20181114
import (
    "context"
    "fmt"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

const APIVersion = "2018-11-14"

type Client struct {
    common.Client
}

func NewClient(credential common.Credentials, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}

func NewGetMetricStatisticsBatchRequest() (request *GetMetricStatisticsBatchRequest) {
    request = &GetMetricStatisticsBatchRequest{
        BaseRequest: &ksyunhttp.BaseRequest{},
    }
    request.Init().WithApiInfo("monitor", APIVersion, "GetMetricStatisticsBatch")
    return
}

func NewGetMetricStatisticsBatchResponse() (response *GetMetricStatisticsBatchResponse) {
    response = &GetMetricStatisticsBatchResponse{
        BaseResponse: &ksyunhttp.BaseResponse{},
    }
    return
}

func (c *Client) GetMetricStatisticsBatch(request *GetMetricStatisticsBatchRequest) string {
    return c.GetMetricStatisticsBatchWithContext(context.Background(), request)
}

func (c *Client) GetMetricStatisticsBatchWithContext(ctx context.Context, request *GetMetricStatisticsBatchRequest) string {
    if request == nil {
        request = NewGetMetricStatisticsBatchRequest()
    }
    request.SetContext(ctx)
    request.SetContentType("application/json")

    response := NewGetMetricStatisticsBatchResponse()
    err, msg := c.Send(request, response)
    if err != nil {
        return fmt.Sprintf("%+v\n", err)
    }
    return msg
}


