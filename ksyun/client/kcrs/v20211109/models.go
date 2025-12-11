package v20211109
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)
type CreateWebhookTriggerTriggerHeader struct {
	Key   *string   `json:"Key,omitempty" name:"Key"`
	Value []*string `json:"Value,omitempty" name:"Value"`
}
type CreateWebhookTriggerTrigger struct {
	TriggerName *string                              `json:"TriggerName,omitempty" name:"TriggerName"`
	TriggerUrl  *string                              `json:"TriggerUrl,omitempty" name:"TriggerUrl"`
	Enabled     *bool                                `json:"Enabled,omitempty" name:"Enabled"`
	EventType   []*string                            `json:"EventType,omitempty" name:"EventType"`
	Header      []*CreateWebhookTriggerTriggerHeader `json:"Header,omitempty" name:"Header"`
}
type CreateRetentionRuleRule struct {
	Scope    *string `json:"Scope,omitempty" name:"Scope"`
	Template *string `json:"Template,omitempty" name:"Template"`
	Tag      *string `json:"Tag,omitempty" name:"Tag"`
	UnTagged *bool   `json:"UnTagged,omitempty" name:"UnTagged"`
	Param    *string `json:"Param,omitempty" name:"Param"`
	Disabled *bool   `json:"Disabled,omitempty" name:"Disabled"`
}
type UpdateRetentionRuleRule struct {
	RuleId   *string `json:"RuleId,omitempty" name:"RuleId"`
	Scope    *string `json:"Scope,omitempty" name:"Scope"`
	Template *string `json:"Template,omitempty" name:"Template"`
	Tag      *string `json:"Tag,omitempty" name:"Tag"`
	UnTagged *bool   `json:"UnTagged,omitempty" name:"UnTagged"`
	Param    *string `json:"Param,omitempty" name:"Param"`
	Disabled *bool   `json:"Disabled,omitempty" name:"Disabled"`
}
type DeleteRetentionRuleRule struct {
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
}


type CreateNamespaceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace   *string `json:"Namespace,omitempty" name:"Namespace"`
	PublicField *string `json:"Public,omitempty" name:"Public"`
}

func (r *CreateNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateNamespaceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeNamespaceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace   *string `json:"Namespace,omitempty" name:"Namespace"`
	PublicField *string `json:"Public,omitempty" name:"Public"`
	MaxResults  *string `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker      *string `json:"Marker,omitempty" name:"Marker"`
}

func (r *DescribeNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeNamespaceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	MaxResults   *int    `json:"MaxResults" name:"MaxResults"`
	Marker       *int    `json:"Marker" name:"Marker"`
	TotalCount   *int    `json:"TotalCount" name:"TotalCount"`
	NamespaceSet []struct {
		Namespace *string `json:"Namespace" name:"Namespace"`
		Public *bool `json:"Public" name:"Public"`
		RepoCount *string `json:"RepoCount" name:"RepoCount"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
	} `json:"NamespaceSet"`
}

func (r *DescribeNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyNamespaceTypeRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace   *string `json:"Namespace,omitempty" name:"Namespace"`
	PublicField *string `json:"Public,omitempty" name:"Public"`
}

func (r *ModifyNamespaceTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyNamespaceTypeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyNamespaceTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNamespaceTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeNamespaceExistRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeNamespaceExistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeNamespaceExistResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Exist     *bool   `json:"Exist" name:"Exist"`
}

func (r *DescribeNamespaceExistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNamespaceExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteNamespaceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DeleteNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteNamespaceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeImagesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	RepoName   *string `json:"RepoName,omitempty" name:"RepoName"`
	ImageId    *string `json:"ImageId,omitempty" name:"ImageId"`
	MaxResults *string `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker     *string `json:"Marker,omitempty" name:"Marker"`
}

func (r *DescribeImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeImagesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	MaxResults *int    `json:"MaxResults" name:"MaxResults"`
	Marker     *int    `json:"Marker" name:"Marker"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	ImageSet   []struct {
		ImageId    *string   `json:"ImageId" name:"ImageId"`
		Size       *int      `json:"Size" name:"Size"`
		CreateTime *string   `json:"CreateTime" name:"CreateTime"`
		TagNames   []*string `json:"TagNames" name:"TagNames"`
	} `json:"ImageSet"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteImagesRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	RepoName   *string `json:"RepoName,omitempty" name:"RepoName"`
	ImageId    *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DeleteImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteImagesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteRepoTagRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	RepoName   *string `json:"RepoName,omitempty" name:"RepoName"`
	TagName    *string `json:"TagName,omitempty" name:"TagName"`
}

func (r *DeleteRepoTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteRepoTagResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteRepoTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRepoTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeRepositoryRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string   `json:"Namespace,omitempty" name:"Namespace"`
	RepoName   []*string `json:"RepoName,omitempty" name:"RepoName"`
	MaxResults *string   `json:"MaxResults,omitempty" name:"MaxResults"`
	Marker     *string   `json:"Marker,omitempty" name:"Marker"`
}

func (r *DescribeRepositoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeRepositoryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	MaxResults *int    `json:"MaxResults" name:"MaxResults"`
	Marker     *int    `json:"Marker" name:"Marker"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	RepoSet    []struct {
		RepoName *string `json:"RepoName" name:"RepoName"`
		Public *bool   `json:"Public" name:"Public"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
		Desc   *string `json:"Desc" name:"Desc"`
	} `json:"RepoSet"`
}

func (r *DescribeRepositoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyRepoDescRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	RepoName   *string `json:"RepoName,omitempty" name:"RepoName"`
	Desc       *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *ModifyRepoDescRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyRepoDescResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyRepoDescResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRepoDescResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteRepositoryRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	RepoName   *string `json:"RepoName,omitempty" name:"RepoName"`
}

func (r *DeleteRepositoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteRepositoryResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteRepositoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRepositoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type StartImageScanRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	RepoName   *string `json:"RepoName,omitempty" name:"RepoName"`
	ImageId    *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *StartImageScanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type StartImageScanResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *StartImageScanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartImageScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeImageScanRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	RepoName   *string `json:"RepoName,omitempty" name:"RepoName"`
	ImageId    *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *DescribeImageScanRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeImageScanResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Status     *string `json:"Status" name:"Status"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	FinishTime *string `json:"FinishTime" name:"FinishTime"`
	Summary    struct {
		High   *int `json:"High" name:"High"`
		Medium *int `json:"Medium" name:"Medium"`
		Low    *int `json:"Low" name:"Low"`
		Negligible *int `json:"Negligible" name:"Negligible"`
		Unknown *int `json:"Unknown" name:"Unknown"`
	} `json:"Summary"`
	VulnerabilitySet struct {
		CveName     *string `json:"CveName" name:"CveName"`
		CveLink     *string `json:"CveLink" name:"CveLink"`
		Description *string `json:"Description" name:"Description"`
		Severity    *string `json:"Severity" name:"Severity"`
		Feature     *string `json:"Feature" name:"Feature"`
		CurrentVersion *string `json:"CurrentVersion" name:"CurrentVersion"`
		RepairedVersion *string `json:"RepairedVersion" name:"RepairedVersion"`
	} `json:"VulnerabilitySet"`
}

func (r *DescribeImageScanResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImageScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateInstanceTokenRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	TokenType  *string `json:"TokenType,omitempty" name:"TokenType"`
	TokenTime  *string `json:"TokenTime,omitempty" name:"TokenTime"`
	Desc       *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateInstanceTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateInstanceTokenResponse struct {
	*ksyunhttp.BaseResponse
	Username   *string `json:"Username" name:"Username"`
	TokenId    *string `json:"tokenId" name:"tokenId"`
	Token      *string `json:"Token" name:"Token"`
	ExpireTime *string `json:"ExpireTime" name:"ExpireTime"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateInstanceTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeInternalEndpointRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInternalEndpointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInternalEndpointResponse struct {
	*ksyunhttp.BaseResponse
	RequestId    *string `json:"RequestId" name:"RequestId"`
	TotalCount   *int    `json:"TotalCount" name:"TotalCount"`
	AccessVpcSet []struct {
		VpcId  *string `json:"VpcId" name:"VpcId"`
		SubnetId *string `json:"SubnetId" name:"SubnetId"`
		Status *string `json:"Status" name:"Status"`
		EniLBIp *string `json:"EniLBIp" name:"EniLBIp"`
	} `json:"AccessVpcSet"`
}

func (r *DescribeInternalEndpointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeInstanceTokenRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Marker     *string `json:"Marker,omitempty" name:"Marker"`
	MaxResults *string `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeInstanceTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceTokenResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	Marker     *int    `json:"Marker" name:"Marker"`
	MaxResults *int    `json:"MaxResults" name:"MaxResults"`
	TokenSet   []struct {
		TokenId *string `json:"TokenId" name:"TokenId"`
		Enable *bool   `json:"Enable" name:"Enable"`
		Desc   *string `json:"Desc" name:"Desc"`
		CreateTime *string `json:"CreateTime" name:"CreateTime"`
		ExpireTime *string `json:"ExpireTime" name:"ExpireTime"`
	} `json:"TokenSet"`
}

func (r *DescribeInstanceTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateInternalEndpointRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId      *string `json:"InstanceId,omitempty" name:"InstanceId"`
	VpcId           *string `json:"VpcId,omitempty" name:"VpcId"`
	ReserveSubnetId *string `json:"ReserveSubnetId,omitempty" name:"ReserveSubnetId"`
}

func (r *CreateInternalEndpointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateInternalEndpointResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateInternalEndpointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInternalEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyInstanceTokenStatusRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	TokenId    *string `json:"TokenId,omitempty" name:"TokenId"`
	Enable     *string `json:"Enable,omitempty" name:"Enable"`
}

func (r *ModifyInstanceTokenStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyInstanceTokenStatusResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyInstanceTokenStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceTokenStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteInternalEndpointRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	VpcId      *string `json:"VpcId,omitempty" name:"VpcId"`
	EniLBIp    *string `json:"EniLBIp,omitempty" name:"EniLBIp"`
}

func (r *DeleteInternalEndpointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteInternalEndpointResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteInternalEndpointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInternalEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyInstanceTokenInformationRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	TokenId    *string `json:"TokenId,omitempty" name:"TokenId"`
	TokenType  *string `json:"TokenType,omitempty" name:"TokenType"`
	TokenTime  *string `json:"TokenTime,omitempty" name:"TokenTime"`
	Desc       *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *ModifyInstanceTokenInformationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyInstanceTokenInformationResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyInstanceTokenInformationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceTokenInformationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeInternalEndpointDnsRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId          *string `json:"InstanceId,omitempty" name:"InstanceId"`
	VpcId               *string `json:"VpcId,omitempty" name:"VpcId"`
	EniLBIp             *string `json:"EniLBIp,omitempty" name:"EniLBIp"`
	InternalEndpointDns *string `json:"InternalEndpointDns,omitempty" name:"InternalEndpointDns"`
}

func (r *DescribeInternalEndpointDnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInternalEndpointDnsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	InternalEndpointDnsSet []struct {
		InternalEndpointDns *string `json:"InternalEndpointDns" name:"InternalEndpointDns"`
		Status *string `json:"Status" name:"Status"`
	} `json:"InternalEndpointDnsSet"`
}

func (r *DescribeInternalEndpointDnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalEndpointDnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteInstanceTokenRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	TokenId    *string `json:"TokenId,omitempty" name:"TokenId"`
}

func (r *DeleteInstanceTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteInstanceTokenResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteInstanceTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateInternalEndpointDnsRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId          *string `json:"InstanceId,omitempty" name:"InstanceId"`
	VpcId               *string `json:"VpcId,omitempty" name:"VpcId"`
	EniLBIp             *string `json:"EniLBIp,omitempty" name:"EniLBIp"`
	InternalEndpointDns *string `json:"InternalEndpointDns,omitempty" name:"InternalEndpointDns"`
}

func (r *CreateInternalEndpointDnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateInternalEndpointDnsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateInternalEndpointDnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInternalEndpointDnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteInternalEndpointDnsRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId          *string `json:"InstanceId,omitempty" name:"InstanceId"`
	VpcId               *string `json:"VpcId,omitempty" name:"VpcId"`
	EniLBIp             *string `json:"EniLBIp,omitempty" name:"EniLBIp"`
	InternalEndpointDns *string `json:"InternalEndpointDns,omitempty" name:"InternalEndpointDns"`
}

func (r *DeleteInternalEndpointDnsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteInternalEndpointDnsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteInternalEndpointDnsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInternalEndpointDnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	ChargeType   *string `json:"ChargeType,omitempty" name:"ChargeType"`
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	PurchaseTime *string `json:"PurchaseTime,omitempty" name:"PurchaseTime"`
	ProjectId    *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	InstanceId *string `json:"InstanceId" name:"InstanceId"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   *string `json:"InstanceId,omitempty" name:"InstanceId"`
	DeleteBucket *bool   `json:"DeleteBucket,omitempty" name:"DeleteBucket"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteInstanceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeInstanceUsageRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceUsageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceUsageResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	NamespaceQuota *int    `json:"NamespaceQuota" name:"NamespaceQuota"`
	NamespaceUsage *int    `json:"NamespaceUsage" name:"NamespaceUsage"`
	RepoQuota      *int    `json:"RepoQuota" name:"RepoQuota"`
	RepoUsage      *int    `json:"RepoUsage" name:"RepoUsage"`
}

func (r *DescribeInstanceUsageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeInstanceRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId   []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	Marker       *string   `json:"Marker,omitempty" name:"Marker"`
	MaxResults   *string   `json:"MaxResults,omitempty" name:"MaxResults"`
	ProjectId    []*string `json:"ProjectId,omitempty" name:"ProjectId"`
	InstanceName *string   `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeInstanceResponse struct {
	*ksyunhttp.BaseResponse
	TotalCount  *int `json:"TotalCount" name:"TotalCount"`
	InstanceSet []struct {
		InstanceId   *string `json:"InstanceId" name:"InstanceId"`
		InstanceName *string `json:"InstanceName" name:"InstanceName"`
		InstanceType *string `json:"InstanceType" name:"InstanceType"`
		InstanceStatus *string `json:"InstanceStatus" name:"InstanceStatus"`
		InternalEndpoint *string `json:"InternalEndpoint" name:"InternalEndpoint"`
		CreateTime   *string `json:"CreateTime" name:"CreateTime"`
		ExpiredTime  *string `json:"ExpiredTime" name:"ExpiredTime"`
		ChargeType   *string `json:"ChargeType" name:"ChargeType"`
		ProjectId    *string `json:"ProjectId" name:"ProjectId"`
	} `json:"InstanceSet"`
	RequestId  *string `json:"RequestId" name:"RequestId"`
	MaxResults *int    `json:"MaxResults" name:"MaxResults"`
	Marker     *int    `json:"Marker" name:"Marker"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateWebhookTriggerRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string                      `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string                      `json:"Namespace,omitempty" name:"Namespace"`
	Trigger    *CreateWebhookTriggerTrigger `json:"Trigger,omitempty" name:"Trigger"`
}

func (r *CreateWebhookTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateWebhookTriggerResponse struct {
	*ksyunhttp.BaseResponse
	TriggerId *int    `json:"triggerId" name:"triggerId"`
	RequestId *string `json:"requestId" name:"requestId"`
}

func (r *CreateWebhookTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeWebhookTriggerRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	TriggerId  *string `json:"TriggerId,omitempty" name:"TriggerId"`
	Marker     *string `json:"Marker,omitempty" name:"Marker"`
	MaxResults *string `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeWebhookTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeWebhookTriggerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	Marker     *int    `json:"Marker" name:"Marker"`
	MaxResults *int    `json:"MaxResults" name:"MaxResults"`
	TriggerSet []struct {
		TriggerId   *int      `json:"TriggerId" name:"TriggerId"`
		TriggerName *string   `json:"TriggerName" name:"TriggerName"`
		EventType   []*string `json:"EventType" name:"EventType"`
		TriggerUrl  *string   `json:"TriggerUrl" name:"TriggerUrl"`
		Enabled     *bool     `json:"Enabled" name:"Enabled"`
		CreateTime  *string   `json:"CreateTime" name:"CreateTime"`
		UpdateTime  *string   `json:"UpdateTime" name:"UpdateTime"`
	} `json:"TriggerSet"`
}

func (r *DescribeWebhookTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type ModifyWebhookTriggerRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  *string   `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace   *string   `json:"Namespace,omitempty" name:"Namespace"`
	Trigger     *string   `json:"Trigger,omitempty" name:"Trigger"`
	TriggerId   *string   `json:"TriggerId,omitempty" name:"TriggerId"`
	TriggerName *string   `json:"TriggerName,omitempty" name:"TriggerName"`
	EventType   []*string `json:"EventType,omitempty" name:"EventType"`
	TriggerUrl  *string   `json:"TriggerUrl,omitempty" name:"TriggerUrl"`
	Header      []*string `json:"Header,omitempty" name:"Header"`
	Enabled     *string   `json:"Enabled,omitempty" name:"Enabled"`
	Key         *string   `json:"Key,omitempty" name:"Key"`
	Value       []*string `json:"Value,omitempty" name:"Value"`
}

func (r *ModifyWebhookTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyWebhookTriggerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ModifyWebhookTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeWebhookTriggerLogRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	TriggerId  *string `json:"TriggerId,omitempty" name:"TriggerId"`
	Marker     *string `json:"Marker,omitempty" name:"Marker"`
	MaxResults *string `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *DescribeWebhookTriggerLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeWebhookTriggerLogResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	MaxResults *int    `json:"MaxResults" name:"MaxResults"`
	Marker     *int    `json:"Marker" name:"Marker"`
	TotalCount *int    `json:"TotalCount" name:"TotalCount"`
	Logs       struct {
	} `json:"Logs"`
}

func (r *DescribeWebhookTriggerLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebhookTriggerLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteWebhookTriggerRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	TriggerId  *string `json:"TriggerId,omitempty" name:"TriggerId"`
}

func (r *DeleteWebhookTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteWebhookTriggerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteWebhookTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWebhookTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type CreateRetentionRuleRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string                  `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string                  `json:"Namespace,omitempty" name:"Namespace"`
	Rule       *CreateRetentionRuleRule `json:"Rule,omitempty" name:"Rule"`
}

func (r *CreateRetentionRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateRetentionRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *CreateRetentionRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRetentionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UpdateRetentionRuleRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string                  `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string                  `json:"Namespace,omitempty" name:"Namespace"`
	Rule       *UpdateRetentionRuleRule `json:"Rule,omitempty" name:"Rule"`
}

func (r *UpdateRetentionRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateRetentionRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *UpdateRetentionRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRetentionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteRetentionRuleRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string                  `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string                  `json:"Namespace,omitempty" name:"Namespace"`
	Rule       *DeleteRetentionRuleRule `json:"Rule,omitempty" name:"Rule"`
}

func (r *DeleteRetentionRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteRetentionRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *DeleteRetentionRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRetentionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DescribeRetentionRuleRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeRetentionRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeRetentionRuleResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Id       *int    `json:"Id" name:"Id"`
		Disabled *bool   `json:"Disabled" name:"Disabled"`
		RegistryScope *string `json:"RegistryScope" name:"RegistryScope"`
		Type     *string `json:"Type" name:"Type"`
		UnTagged *bool   `json:"UnTagged" name:"UnTagged"`
		Template *string `json:"Template" name:"Template"`
		TagPatten *string `json:"TagPatten" name:"TagPatten"`
		Param    *int    `json:"Param" name:"Param"`
	} `json:"Data"`
}

func (r *DescribeRetentionRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRetentionRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type RunRetentionPolicyRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	TestRun    *bool   `json:"TestRun,omitempty" name:"TestRun"`
}

func (r *RunRetentionPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RunRetentionPolicyResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *RunRetentionPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunRetentionPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type GetRetentionPolicyLogsRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	Page       *int    `json:"Page,omitempty" name:"Page"`
	PageSize   *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *GetRetentionPolicyLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRetentionPolicyLogsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Id       *string `json:"Id" name:"Id"`
		EndTime  *string `json:"EndTime" name:"EndTime"`
		StartTime *string `json:"StartTime" name:"StartTime"`
		Status   *string `json:"Status" name:"Status"`
		Trigger  *string `json:"Trigger" name:"Trigger"`
		DryRun   *bool   `json:"DryRun" name:"DryRun"`
		TakeTime *int    `json:"TakeTime" name:"TakeTime"`
	} `json:"Data"`
	Page     *int `json:"Page" name:"Page"`
	PageSize *int `json:"PageSize" name:"PageSize"`
	Total    *int `json:"Total" name:"Total"`
}

func (r *GetRetentionPolicyLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRetentionPolicyLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type GetRetentionPolicyLogDetailRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace   *string `json:"Namespace,omitempty" name:"Namespace"`
	ExecutionId *string `json:"ExecutionId,omitempty" name:"ExecutionId"`
	Page        *int    `json:"Page,omitempty" name:"Page"`
	PageSize    *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *GetRetentionPolicyLogDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRetentionPolicyLogDetailResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		LogDetail struct {
			ExecutionId *string `json:"ExecutionId" name:"ExecutionId"`
			Id          *int    `json:"Id" name:"Id"`
			EndTime     *string `json:"EndTime" name:"EndTime"`
			StartTime   *string `json:"StartTime" name:"StartTime"`
			Status      *string `json:"Status" name:"Status"`
			StatusCode  *string `json:"StatusCode" name:"StatusCode"`
			Repository  *string `json:"Repository" name:"Repository"`
			Rentained   *int    `json:"Rentained" name:"Rentained"`
			Total       *int    `json:"Total" name:"Total"`
			TakeTime    *int    `json:"TakeTime" name:"TakeTime"`
		} `json:"LogDetail" name:"LogDetail"`
		Page *int `json:"Page" name:"Page"`
		PageSize *int `json:"PageSize" name:"PageSize"`
		Total *int `json:"Total" name:"Total"`
	} `json:"Data"`
}

func (r *GetRetentionPolicyLogDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRetentionPolicyLogDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type GetRetentionPolicyLogRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId  *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace   *string `json:"Namespace,omitempty" name:"Namespace"`
	TaskId      *int    `json:"TaskId,omitempty" name:"TaskId"`
	ExecutionId *int    `json:"ExecutionId,omitempty" name:"ExecutionId"`
}

func (r *GetRetentionPolicyLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRetentionPolicyLogResponse struct {
	*ksyunhttp.BaseResponse
	Data *string `json:"Data" name:"Data"`
}

func (r *GetRetentionPolicyLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRetentionPolicyLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type GetRetentionTriggerRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *GetRetentionTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetRetentionTriggerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		DisplayText *string   `json:"DisplayText" name:"DisplayText"`
		Cron        *string   `json:"Cron" name:"Cron"`
		Optional    []*string `json:"Optional" name:"Optional"`
	} `json:"Data"`
}

func (r *GetRetentionTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRetentionTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type UpdateRetentionTriggerRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	Namespace  *string `json:"Namespace,omitempty" name:"Namespace"`
	Trigger    *string `json:"Trigger,omitempty" name:"Trigger"`
}

func (r *UpdateRetentionTriggerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateRetentionTriggerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *UpdateRetentionTriggerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateRetentionTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScheduleRequest struct {
	*ksyunhttp.BaseRequest
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	TestRun    *bool   `json:"TestRun,omitempty" name:"TestRun"`
}

func (r *ScheduleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ScheduleResponse struct {
	*ksyunhttp.BaseResponse
	Requestid *string `json:"Requestid" name:"Requestid"`
	Data      *string `json:"Data" name:"Data"`
}

func (r *ScheduleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

