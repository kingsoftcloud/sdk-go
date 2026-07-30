package v20251114

import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)

type RetrieveKnowledgeRetrievalModelRerankingMode struct {
	RerankingProviderName *string `json:"RerankingProviderName,omitempty" name:"RerankingProviderName"`
	RerankingModelName    *string `json:"RerankingModelName,omitempty" name:"RerankingModelName"`
}
type RetrieveKnowledgeRetrievalModelRetrieverVector struct {
	TopK                  *int     `json:"TopK,omitempty" name:"TopK"`
	ScoreThreshold        *float64 `json:"ScoreThreshold,omitempty" name:"ScoreThreshold"`
	ScoreThresholdEnabled *bool    `json:"ScoreThresholdEnabled,omitempty" name:"ScoreThresholdEnabled"`
}
type RetrieveKnowledgeRetrievalModelRetrieverInverted struct {
	TopK                  *int     `json:"TopK,omitempty" name:"TopK"`
	ScoreThreshold        *float64 `json:"ScoreThreshold,omitempty" name:"ScoreThreshold"`
	ScoreThresholdEnabled *bool    `json:"ScoreThresholdEnabled,omitempty" name:"ScoreThresholdEnabled"`
}
type RetrieveKnowledgeRetrievalModelRetriever struct {
	Vector   *RetrieveKnowledgeRetrievalModelRetrieverVector   `json:"Vector,omitempty" name:"Vector"`
	Inverted *RetrieveKnowledgeRetrievalModelRetrieverInverted `json:"Inverted,omitempty" name:"Inverted"`
}
type RetrieveKnowledgeRetrievalModel struct {
	SearchMethod          *string                                       `json:"SearchMethod,omitempty" name:"SearchMethod"`
	RerankingEnable       *bool                                         `json:"RerankingEnable,omitempty" name:"RerankingEnable"`
	RerankingMode         *RetrieveKnowledgeRetrievalModelRerankingMode `json:"RerankingMode,omitempty" name:"RerankingMode"`
	TopK                  *int                                          `json:"TopK,omitempty" name:"TopK"`
	ScoreThresholdEnabled *bool                                         `json:"ScoreThresholdEnabled,omitempty" name:"ScoreThresholdEnabled"`
	ScoreThreshold        *float64                                      `json:"ScoreThreshold,omitempty" name:"ScoreThreshold"`
	Retriever             *RetrieveKnowledgeRetrievalModelRetriever     `json:"Retriever,omitempty" name:"Retriever"`
}
type ImportDocumentsDataProcessRuleRulesSegmentation struct {
	Delimiter *string `json:"Delimiter,omitempty" name:"Delimiter"`
	MaxWords  *int    `json:"MaxWords,omitempty" name:"MaxWords"`
}
type ImportDocumentsDataProcessRuleRules struct {
	Segmentation *ImportDocumentsDataProcessRuleRulesSegmentation `json:"Segmentation,omitempty" name:"Segmentation"`
}
type ImportDocumentsDataProcessRule struct {
	Mode  *string                              `json:"Mode,omitempty" name:"Mode"`
	Rules *ImportDocumentsDataProcessRuleRules `json:"Rules,omitempty" name:"Rules"`
}
type ImportDocumentsDataSchemaColumns struct {
	OriginalColumnName *string `json:"OriginalColumnName,omitempty" name:"OriginalColumnName"`
	ModifiedColumnName *string `json:"ModifiedColumnName,omitempty" name:"ModifiedColumnName"`
	Description        *string `json:"Description,omitempty" name:"Description"`
	IsIndex            *bool   `json:"IsIndex,omitempty" name:"IsIndex"`
	CType              *string `json:"CType,omitempty" name:"CType"`
}
type ImportDocumentsDataSchema struct {
	TableName   *string                             `json:"TableName,omitempty" name:"TableName"`
	Description *string                             `json:"Description,omitempty" name:"Description"`
	Columns     []*ImportDocumentsDataSchemaColumns `json:"Columns,omitempty" name:"Columns"`
}
type ImportDocumentsDataAnalysisStrategyImageUnderstandingVlmPrompt struct {
	Role *string `json:"Role,omitempty" name:"Role"`
	Text *string `json:"Text,omitempty" name:"Text"`
}
type ImportDocumentsDataAnalysisStrategyImageUnderstandingVlm struct {
	ModelName    *string                                                           `json:"ModelName,omitempty" name:"ModelName"`
	ProviderName *string                                                           `json:"ProviderName,omitempty" name:"ProviderName"`
	Prompt       []*ImportDocumentsDataAnalysisStrategyImageUnderstandingVlmPrompt `json:"Prompt,omitempty" name:"Prompt"`
}
type ImportDocumentsDataAnalysisStrategyImageUnderstanding struct {
	Enable *bool                                                     `json:"Enable,omitempty" name:"Enable"`
	Method *string                                                   `json:"Method,omitempty" name:"Method"`
	Vlm    *ImportDocumentsDataAnalysisStrategyImageUnderstandingVlm `json:"Vlm,omitempty" name:"Vlm"`
}
type ImportDocumentsDataAnalysisStrategyTableUnderstandingVlmPrompt struct {
	Role *string `json:"Role,omitempty" name:"Role"`
	Text *string `json:"Text,omitempty" name:"Text"`
}
type ImportDocumentsDataAnalysisStrategyTableUnderstandingVlm struct {
	ModelName    *string                                                           `json:"ModelName,omitempty" name:"ModelName"`
	ProviderName *string                                                           `json:"ProviderName,omitempty" name:"ProviderName"`
	Prompt       []*ImportDocumentsDataAnalysisStrategyTableUnderstandingVlmPrompt `json:"Prompt,omitempty" name:"Prompt"`
}
type ImportDocumentsDataAnalysisStrategyTableUnderstanding struct {
	Enable *bool                                                     `json:"Enable,omitempty" name:"Enable"`
	Method *string                                                   `json:"Method,omitempty" name:"Method"`
	Vlm    *ImportDocumentsDataAnalysisStrategyTableUnderstandingVlm `json:"Vlm,omitempty" name:"Vlm"`
}
type ImportDocumentsDataAnalysisStrategy struct {
	ImageUnderstanding *ImportDocumentsDataAnalysisStrategyImageUnderstanding `json:"ImageUnderstanding,omitempty" name:"ImageUnderstanding"`
	TableUnderstanding *ImportDocumentsDataAnalysisStrategyTableUnderstanding `json:"TableUnderstanding,omitempty" name:"TableUnderstanding"`
}
type ImportDocumentsDataRetrievalModelRerankingMode struct {
	RerankingProviderName *string `json:"RerankingProviderName,omitempty" name:"RerankingProviderName"`
	RerankingModelName    *string `json:"RerankingModelName,omitempty" name:"RerankingModelName"`
}
type ImportDocumentsDataRetrievalModelRetrieverVector struct {
	TopK                  *int     `json:"TopK,omitempty" name:"TopK"`
	ScoreThreshold        *float64 `json:"ScoreThreshold,omitempty" name:"ScoreThreshold"`
	ScoreThresholdEnabled *bool    `json:"ScoreThresholdEnabled,omitempty" name:"ScoreThresholdEnabled"`
}
type ImportDocumentsDataRetrievalModelRetrieverInverted struct {
	TopK                  *int     `json:"TopK,omitempty" name:"TopK"`
	ScoreThreshold        *float64 `json:"ScoreThreshold,omitempty" name:"ScoreThreshold"`
	ScoreThresholdEnabled *bool    `json:"ScoreThresholdEnabled,omitempty" name:"ScoreThresholdEnabled"`
}
type ImportDocumentsDataRetrievalModelRetriever struct {
	Vector   *ImportDocumentsDataRetrievalModelRetrieverVector   `json:"Vector,omitempty" name:"Vector"`
	Inverted *ImportDocumentsDataRetrievalModelRetrieverInverted `json:"Inverted,omitempty" name:"Inverted"`
}
type ImportDocumentsDataRetrievalModel struct {
	SearchMethod          *string                                         `json:"SearchMethod,omitempty" name:"SearchMethod"`
	RerankingEnable       *bool                                           `json:"RerankingEnable,omitempty" name:"RerankingEnable"`
	RerankingMode         *ImportDocumentsDataRetrievalModelRerankingMode `json:"RerankingMode,omitempty" name:"RerankingMode"`
	TopK                  *int                                            `json:"TopK,omitempty" name:"TopK"`
	ScoreThresholdEnabled *bool                                           `json:"ScoreThresholdEnabled,omitempty" name:"ScoreThresholdEnabled"`
	ScoreThreshold        *float64                                        `json:"ScoreThreshold,omitempty" name:"ScoreThreshold"`
	Retriever             *ImportDocumentsDataRetrievalModelRetriever     `json:"Retriever,omitempty" name:"Retriever"`
}
type ImportDocumentsData struct {
	IndexingTechnique *string                              `json:"IndexingTechnique,omitempty" name:"IndexingTechnique"`
	ProcessRule       *ImportDocumentsDataProcessRule      `json:"ProcessRule,omitempty" name:"ProcessRule"`
	DataType          *string                              `json:"DataType,omitempty" name:"DataType"`
	Schema            *ImportDocumentsDataSchema           `json:"Schema,omitempty" name:"Schema"`
	AnalysisStrategy  *ImportDocumentsDataAnalysisStrategy `json:"AnalysisStrategy,omitempty" name:"AnalysisStrategy"`
	RetrievalModel    *ImportDocumentsDataRetrievalModel   `json:"RetrievalModel,omitempty" name:"RetrievalModel"`
}
type ModifyKnowledgeBaseRetrievalModelRerankingMode struct {
	RerankingProviderName *string `json:"RerankingProviderName,omitempty" name:"RerankingProviderName"`
	RerankingModelName    *string `json:"RerankingModelName,omitempty" name:"RerankingModelName"`
}
type ModifyKnowledgeBaseRetrievalModelRetrieverVector struct {
	TopK                  *int     `json:"TopK,omitempty" name:"TopK"`
	ScoreThreshold        *float64 `json:"ScoreThreshold,omitempty" name:"ScoreThreshold"`
	ScoreThresholdEnabled *bool    `json:"ScoreThresholdEnabled,omitempty" name:"ScoreThresholdEnabled"`
}
type ModifyKnowledgeBaseRetrievalModelRetrieverInverted struct {
	TopK                  *int     `json:"TopK,omitempty" name:"TopK"`
	ScoreThreshold        *float64 `json:"ScoreThreshold,omitempty" name:"ScoreThreshold"`
	ScoreThresholdEnabled *bool    `json:"ScoreThresholdEnabled,omitempty" name:"ScoreThresholdEnabled"`
}
type ModifyKnowledgeBaseRetrievalModelRetriever struct {
	Vector   *ModifyKnowledgeBaseRetrievalModelRetrieverVector   `json:"Vector,omitempty" name:"Vector"`
	Inverted *ModifyKnowledgeBaseRetrievalModelRetrieverInverted `json:"Inverted,omitempty" name:"Inverted"`
}
type ModifyKnowledgeBaseRetrievalModel struct {
	SearchMethod          *string                                         `json:"SearchMethod,omitempty" name:"SearchMethod"`
	RerankingEnable       *bool                                           `json:"RerankingEnable,omitempty" name:"RerankingEnable"`
	RerankingMode         *ModifyKnowledgeBaseRetrievalModelRerankingMode `json:"RerankingMode,omitempty" name:"RerankingMode"`
	TopK                  *int                                            `json:"TopK,omitempty" name:"TopK"`
	ScoreThresholdEnabled *bool                                           `json:"ScoreThresholdEnabled,omitempty" name:"ScoreThresholdEnabled"`
	ScoreThreshold        *float64                                        `json:"ScoreThreshold,omitempty" name:"ScoreThreshold"`
	Retriever             *ModifyKnowledgeBaseRetrievalModelRetriever     `json:"Retriever,omitempty" name:"Retriever"`
}
type CreateKnowledgeBaseRetrievalModelRerankingMode struct {
	RerankingProviderName *string `json:"RerankingProviderName,omitempty" name:"RerankingProviderName"`
	RerankingModelName    *string `json:"RerankingModelName,omitempty" name:"RerankingModelName"`
}
type CreateKnowledgeBaseRetrievalModelRetrieverVector struct {
	TopK                  *int     `json:"TopK,omitempty" name:"TopK"`
	ScoreThreshold        *float64 `json:"ScoreThreshold,omitempty" name:"ScoreThreshold"`
	ScoreThresholdEnabled *bool    `json:"ScoreThresholdEnabled,omitempty" name:"ScoreThresholdEnabled"`
}
type CreateKnowledgeBaseRetrievalModelRetrieverInverted struct {
	TopK                  *int     `json:"TopK,omitempty" name:"TopK"`
	ScoreThreshold        *float64 `json:"ScoreThreshold,omitempty" name:"ScoreThreshold"`
	ScoreThresholdEnabled *bool    `json:"ScoreThresholdEnabled,omitempty" name:"ScoreThresholdEnabled"`
}
type CreateKnowledgeBaseRetrievalModelRetriever struct {
	Vector   *CreateKnowledgeBaseRetrievalModelRetrieverVector   `json:"Vector,omitempty" name:"Vector"`
	Inverted *CreateKnowledgeBaseRetrievalModelRetrieverInverted `json:"Inverted,omitempty" name:"Inverted"`
}
type CreateKnowledgeBaseRetrievalModel struct {
	SearchMethod          *string                                         `json:"SearchMethod,omitempty" name:"SearchMethod"`
	RerankingEnable       *bool                                           `json:"RerankingEnable,omitempty" name:"RerankingEnable"`
	RerankingMode         *CreateKnowledgeBaseRetrievalModelRerankingMode `json:"RerankingMode,omitempty" name:"RerankingMode"`
	TopK                  *int                                            `json:"TopK,omitempty" name:"TopK"`
	ScoreThresholdEnabled *bool                                           `json:"ScoreThresholdEnabled,omitempty" name:"ScoreThresholdEnabled"`
	ScoreThreshold        *float64                                        `json:"ScoreThreshold,omitempty" name:"ScoreThreshold"`
	Retriever             *CreateKnowledgeBaseRetrievalModelRetriever     `json:"Retriever,omitempty" name:"Retriever"`
}
type CreateMemorySdkDataConversationContent struct {
	Type *string `json:"Type,omitempty" name:"Type"`
	Text *string `json:"Text,omitempty" name:"Text"`
}
type CreateMemorySdkDataConversation struct {
	Role      *string                                   `json:"Role,omitempty" name:"Role"`
	CreatedAt *int64                                    `json:"CreatedAt,omitempty" name:"CreatedAt"`
	MessageId *string                                   `json:"MessageId,omitempty" name:"MessageId"`
	Content   []*CreateMemorySdkDataConversationContent `json:"Content,omitempty" name:"Content"`
}
type CreateMemorySdkData struct {
	Conversation []*CreateMemorySdkDataConversation `json:"Conversation,omitempty" name:"Conversation"`
}
type CreateMemoryCollectionLongTermConfigurationStrategies struct {
	Type *string `json:"Type,omitempty" name:"Type"`
}
type CreateMemoryCollectionLongTermConfiguration struct {
	Strategies []*CreateMemoryCollectionLongTermConfigurationStrategies `json:"Strategies,omitempty" name:"Strategies"`
}
type UpdateMemoryCollectionLongTermConfigurationStrategies struct {
	Type *string `json:"Type,omitempty" name:"Type"`
}
type UpdateMemoryCollectionLongTermConfiguration struct {
	Strategies []*UpdateMemoryCollectionLongTermConfigurationStrategies `json:"Strategies,omitempty" name:"Strategies"`
}
type CreateMcpServerMcpRuntimeConfigKs3Config struct {
	Path *string `json:"Path,omitempty" name:"Path"`
}
type CreateMcpServerMcpRuntimeConfigImagesConfig struct {
	ImageType              *string `json:"ImageType,omitempty" name:"ImageType"`
	NameSpace              *string `json:"NameSpace,omitempty" name:"NameSpace"`
	ImageRepo              *string `json:"ImageRepo,omitempty" name:"ImageRepo"`
	ImageVersion           *string `json:"ImageVersion,omitempty" name:"ImageVersion"`
	EnterpriseInstanceId   *string `json:"EnterpriseInstanceId,omitempty" name:"EnterpriseInstanceId"`
	EnterpriseInstanceName *string `json:"EnterpriseInstanceName,omitempty" name:"EnterpriseInstanceName"`
	UserName               *string `json:"UserName,omitempty" name:"UserName"`
	Password               *string `json:"Password,omitempty" name:"Password"`
}
type CreateMcpServerMcpRuntimeConfigResource struct {
	Cpu    *int `json:"Cpu,omitempty" name:"Cpu"`
	Memory *int `json:"Memory,omitempty" name:"Memory"`
}
type CreateMcpServerMcpRuntimeConfigAdvancedEnvironmentVariables struct {
	Key   *string `json:"Key,omitempty" name:"Key"`
	Value *string `json:"Value,omitempty" name:"Value"`
}
type CreateMcpServerMcpRuntimeConfigAdvanced struct {
	EnvironmentVariables []*CreateMcpServerMcpRuntimeConfigAdvancedEnvironmentVariables `json:"EnvironmentVariables,omitempty" name:"EnvironmentVariables"`
}
type CreateMcpServerMcpRuntimeConfig struct {
	CodeFrom     *string                                      `json:"CodeFrom,omitempty" name:"CodeFrom"`
	Ks3Config    *CreateMcpServerMcpRuntimeConfigKs3Config    `json:"Ks3Config,omitempty" name:"Ks3Config"`
	ImagesConfig *CreateMcpServerMcpRuntimeConfigImagesConfig `json:"ImagesConfig,omitempty" name:"ImagesConfig"`
	Resource     *CreateMcpServerMcpRuntimeConfigResource     `json:"Resource,omitempty" name:"Resource"`
	Advanced     *CreateMcpServerMcpRuntimeConfigAdvanced     `json:"Advanced,omitempty" name:"Advanced"`
}

type DescribeKnowledgeBaseModelsRequest struct {
	*ksyunhttp.BaseRequest
	ModelType *string `json:"ModelType,omitempty" name:"ModelType"`
}

func (r *DescribeKnowledgeBaseModelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeKnowledgeBaseModelsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Models    []struct {
		ModelName     *string `json:"ModelName" name:"ModelName"`
		ModelProvider *string `json:"ModelProvider" name:"ModelProvider"`
		ModelType     *string `json:"ModelType" name:"ModelType"`
	} `json:"Models"`
}

func (r *DescribeKnowledgeBaseModelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKnowledgeBaseModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActivateKnowledgeBaseServiceRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *ActivateKnowledgeBaseServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ActivateKnowledgeBaseServiceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *ActivateKnowledgeBaseServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateKnowledgeBaseServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetrieveKnowledgeRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId      *string                          `json:"DatasetId,omitempty" name:"DatasetId"`
	Query          *string                          `json:"Query,omitempty" name:"Query"`
	RetrievalModel *RetrieveKnowledgeRetrievalModel `json:"RetrievalModel,omitempty" name:"RetrievalModel"`
}

func (r *RetrieveKnowledgeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RetrieveKnowledgeResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Query     struct {
		Content *string `json:"Content" name:"Content"`
	} `json:"Query"`
	Records []struct {
		Segment struct {
			Id            *string   `json:"Id" name:"Id"`
			Position      *int      `json:"Position" name:"Position"`
			DocumentId    *string   `json:"DocumentId" name:"DocumentId"`
			Content       *string   `json:"Content" name:"Content"`
			Answer        *string   `json:"Answer" name:"Answer"`
			WordCount     *int      `json:"WordCount" name:"WordCount"`
			Tokens        *int      `json:"Tokens" name:"Tokens"`
			Keywords      []*string `json:"Keywords" name:"Keywords"`
			IndexNodeId   *string   `json:"IndexNodeId" name:"IndexNodeId"`
			IndexNodeHash *string   `json:"IndexNodeHash" name:"IndexNodeHash"`
			HitCount      *int      `json:"HitCount" name:"HitCount"`
			Enabled       *bool     `json:"Enabled" name:"Enabled"`
			DisabledAt    *int64    `json:"DisabledAt" name:"DisabledAt"`
			DisabledBy    *string   `json:"DisabledBy" name:"DisabledBy"`
			Status        *string   `json:"Status" name:"Status"`
			CreatedBy     *string   `json:"CreatedBy" name:"CreatedBy"`
			CreatedAt     *int64    `json:"CreatedAt" name:"CreatedAt"`
			IndexingAt    *int64    `json:"IndexingAt" name:"IndexingAt"`
			CompletedAt   *int64    `json:"CompletedAt" name:"CompletedAt"`
			Error         *string   `json:"Error" name:"Error"`
			StoppedAt     *int64    `json:"StoppedAt" name:"StoppedAt"`
			Document      struct {
				Id             *string `json:"Id" name:"Id"`
				DataSourceType *string `json:"DataSourceType" name:"DataSourceType"`
				Name           *string `json:"Name" name:"Name"`
			} `json:"Document"`
		} `json:"Segment" name:"Segment"`
		Score        *float64 `json:"Score" name:"Score"`
		TsnePosition *string  `json:"TsnePosition" name:"TsnePosition"`
	} `json:"Records"`
}

func (r *RetrieveKnowledgeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetrieveKnowledgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeChunkRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId  *string `json:"DatasetId,omitempty" name:"DatasetId"`
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

func (r *DescribeChunkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeChunkResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Id        *string `json:"Id" name:"Id"`
		Content   *string `json:"Content" name:"Content"`
		Position  *int    `json:"Position" name:"Position"`
		WordCount *int    `json:"WordCount" name:"WordCount"`
		Tokens    *int    `json:"Tokens" name:"Tokens"`
	} `json:"Data"`
}

func (r *DescribeChunkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeChunkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchDisplayStatusRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId   *string   `json:"DatasetId,omitempty" name:"DatasetId"`
	DocumentIds []*string `json:"DocumentIds,omitempty" name:"DocumentIds"`
}

func (r *BatchDisplayStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type BatchDisplayStatusResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Status    *string `json:"Status" name:"Status"`
	Data      []struct {
		Id            *string `json:"Id" name:"Id"`
		DisplayStatus *string `json:"DisplayStatus" name:"DisplayStatus"`
	} `json:"Data"`
	Errors []struct {
		Id    *string `json:"Id" name:"Id"`
		Error *string `json:"Error" name:"Error"`
	} `json:"Errors"`
}

func (r *BatchDisplayStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BatchDisplayStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisplayStatusRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId  *string `json:"DatasetId,omitempty" name:"DatasetId"`
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

func (r *DisplayStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DisplayStatusResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	Id            *string `json:"Id" name:"Id"`
	DisplayStatus *string `json:"DisplayStatus" name:"DisplayStatus"`
}

func (r *DisplayStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisplayStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IndexingStatusRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`
	Batch     *string `json:"Batch,omitempty" name:"Batch"`
}

func (r *IndexingStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type IndexingStatusResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Id                   *string `json:"Id" name:"Id"`
		IndexingStatus       *string `json:"IndexingStatus" name:"IndexingStatus"`
		ProcessingStartedAt  *int64  `json:"ProcessingStartedAt" name:"ProcessingStartedAt"`
		ParsingCompletedAt   *int64  `json:"ParsingCompletedAt" name:"ParsingCompletedAt"`
		CleaningCompletedAt  *int64  `json:"CleaningCompletedAt" name:"CleaningCompletedAt"`
		SplittingCompletedAt *int64  `json:"SplittingCompletedAt" name:"SplittingCompletedAt"`
		CompletedAt          *int64  `json:"CompletedAt" name:"CompletedAt"`
		PausedAt             *int64  `json:"PausedAt" name:"PausedAt"`
		Error                *string `json:"Error" name:"Error"`
		StoppedAt            *int64  `json:"StoppedAt" name:"StoppedAt"`
		CompletedSegments    *int    `json:"CompletedSegments" name:"CompletedSegments"`
		TotalSegments        *int    `json:"TotalSegments" name:"TotalSegments"`
	} `json:"Data"`
}

func (r *IndexingStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IndexingStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDocumentRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId  *string `json:"DatasetId,omitempty" name:"DatasetId"`
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
}

func (r *DeleteDocumentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteDocumentResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *DeleteDocumentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDocumentRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId  *string `json:"DatasetId,omitempty" name:"DatasetId"`
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
	Metadata   *string `json:"Metadata,omitempty" name:"Metadata"`
}

func (r *DescribeDocumentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDocumentResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string `json:"RequestId" name:"RequestId"`
	Id             *string `json:"Id" name:"Id"`
	Position       *int    `json:"Position" name:"Position"`
	DataSourceType *string `json:"DataSourceType" name:"DataSourceType"`
	DataSourceInfo struct {
		UploadFile struct {
			Id   *string `json:"Id" name:"Id"`
			Name *string `json:"Name" name:"Name"`
			Size *int64  `json:"Size" name:"Size"`
		} `json:"UploadFile" name:"UploadFile"`
	} `json:"DataSourceInfo"`
	DatasetProcessRuleId *string `json:"DatasetProcessRuleId" name:"DatasetProcessRuleId"`
	DatasetProcessRule   struct {
		Mode  *string `json:"Mode" name:"Mode"`
		Rules struct {
			Segmentation struct {
				Delimiter *string `json:"Delimiter" name:"Delimiter"`
				MaxWords  *int    `json:"MaxWords" name:"MaxWords"`
			} `json:"Segmentation"`
		} `json:"Rules" name:"Rules"`
	} `json:"DatasetProcessRule"`
	DocumentProcessRule struct {
		Mode  *string `json:"Mode" name:"Mode"`
		Rules struct {
			Segmentation struct {
				Delimiter *string `json:"Delimiter" name:"Delimiter"`
				MaxWords  *int    `json:"MaxWords" name:"MaxWords"`
			} `json:"Segmentation"`
		} `json:"Rules" name:"Rules"`
	} `json:"DocumentProcessRule"`
	Name                 *string  `json:"Name" name:"Name"`
	CreatedAt            *int64   `json:"CreatedAt" name:"CreatedAt"`
	Tokens               *int     `json:"Tokens" name:"Tokens"`
	IndexingStatus       *string  `json:"IndexingStatus" name:"IndexingStatus"`
	CompletedAt          *int64   `json:"CompletedAt" name:"CompletedAt"`
	UpdatedAt            *int64   `json:"UpdatedAt" name:"UpdatedAt"`
	IndexingLatency      *float64 `json:"IndexingLatency" name:"IndexingLatency"`
	Error                *string  `json:"Error" name:"Error"`
	Enabled              *bool    `json:"Enabled" name:"Enabled"`
	DisabledAt           *int64   `json:"DisabledAt" name:"DisabledAt"`
	Archived             *bool    `json:"Archived" name:"Archived"`
	SegmentCount         *int     `json:"SegmentCount" name:"SegmentCount"`
	AverageSegmentLength *int     `json:"AverageSegmentLength" name:"AverageSegmentLength"`
	HitCount             *int     `json:"HitCount" name:"HitCount"`
	DisplayStatus        *string  `json:"DisplayStatus" name:"DisplayStatus"`
	DocForm              *string  `json:"DocForm" name:"DocForm"`
	DocLanguage          *string  `json:"DocLanguage" name:"DocLanguage"`
}

func (r *DescribeDocumentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDocumentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDocumentsRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`
	Keyword   *string `json:"Keyword,omitempty" name:"Keyword"`
	Page      *int    `json:"Page,omitempty" name:"Page"`
	Limit     *int    `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDocumentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeDocumentsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Id            *string `json:"Id" name:"Id"`
		Name          *string `json:"Name" name:"Name"`
		CreatedAt     *int64  `json:"CreatedAt" name:"CreatedAt"`
		DisplayStatus *string `json:"DisplayStatus" name:"DisplayStatus"`
		Error         *string `json:"Error" name:"Error"`
	} `json:"Data"`
	HasMore *bool `json:"HasMore" name:"HasMore"`
	Total   *int  `json:"Total" name:"Total"`
	Page    *int  `json:"Page" name:"Page"`
	Limit   *int  `json:"Limit" name:"Limit"`
}

func (r *DescribeDocumentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDocumentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportDocumentsRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId *string              `json:"DatasetId,omitempty" name:"DatasetId"`
	Data      *ImportDocumentsData `json:"Data,omitempty" name:"Data"`
	AddType   *string              `json:"AddType,omitempty" name:"AddType"`
	Ks3Path   []*string            `json:"Ks3Path,omitempty" name:"Ks3Path"`
}

func (r *ImportDocumentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ImportDocumentsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Documents []struct {
		Document struct {
			Id            *string `json:"Id" name:"Id"`
			Name          *string `json:"Name" name:"Name"`
			Error         *string `json:"Error" name:"Error"`
			DisplayStatus *string `json:"DisplayStatus" name:"DisplayStatus"`
		} `json:"Document" name:"Document"`
		Batch *string `json:"Batch" name:"Batch"`
	} `json:"Documents"`
}

func (r *ImportDocumentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportDocumentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteKnowledgeBaseRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`
}

func (r *DeleteKnowledgeBaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteKnowledgeBaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *DeleteKnowledgeBaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteKnowledgeBaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyKnowledgeBaseRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId              *string                            `json:"DatasetId,omitempty" name:"DatasetId"`
	Name                   *string                            `json:"Name,omitempty" name:"Name"`
	IndexingTechnique      *string                            `json:"IndexingTechnique,omitempty" name:"IndexingTechnique"`
	EmbeddingModelProvider *string                            `json:"EmbeddingModelProvider,omitempty" name:"EmbeddingModelProvider"`
	EmbeddingModel         *string                            `json:"EmbeddingModel,omitempty" name:"EmbeddingModel"`
	RetrievalModel         *ModifyKnowledgeBaseRetrievalModel `json:"RetrievalModel,omitempty" name:"RetrievalModel"`
	ComputeUnit            *int                               `json:"ComputeUnit,omitempty" name:"ComputeUnit"`
}

func (r *ModifyKnowledgeBaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyKnowledgeBaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Id        *string `json:"Id" name:"Id"`
}

func (r *ModifyKnowledgeBaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKnowledgeBaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKnowledgeBaseRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`
}

func (r *DescribeKnowledgeBaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeKnowledgeBaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId              *string `json:"RequestId" name:"RequestId"`
	Id                     *string `json:"Id" name:"Id"`
	Name                   *string `json:"Name" name:"Name"`
	Description            *string `json:"Description" name:"Description"`
	Provider               *string `json:"Provider" name:"Provider"`
	DataSourceType         *string `json:"DataSourceType" name:"DataSourceType"`
	IndexingTechnique      *string `json:"IndexingTechnique" name:"IndexingTechnique"`
	AppCount               *int    `json:"AppCount" name:"AppCount"`
	DocumentCount          *int    `json:"DocumentCount" name:"DocumentCount"`
	CreatedBy              *string `json:"CreatedBy" name:"CreatedBy"`
	CreatedAt              *int64  `json:"CreatedAt" name:"CreatedAt"`
	UpdatedBy              *string `json:"UpdatedBy" name:"UpdatedBy"`
	UpdatedAt              *int64  `json:"UpdatedAt" name:"UpdatedAt"`
	EmbeddingModel         *string `json:"EmbeddingModel" name:"EmbeddingModel"`
	EmbeddingModelProvider *string `json:"EmbeddingModelProvider" name:"EmbeddingModelProvider"`
	EmbeddingAvailable     *bool   `json:"EmbeddingAvailable" name:"EmbeddingAvailable"`
	RetrievalModelDict     struct {
		SearchMethod    *string `json:"SearchMethod" name:"SearchMethod"`
		RerankingEnable *bool   `json:"RerankingEnable" name:"RerankingEnable"`
		RerankingMode   struct {
			RerankingProviderName *string `json:"RerankingProviderName" name:"RerankingProviderName"`
			RerankingModelName    *string `json:"RerankingModelName" name:"RerankingModelName"`
		} `json:"RerankingMode" name:"RerankingMode"`
		TopK                  *int     `json:"TopK" name:"TopK"`
		ScoreThresholdEnabled *bool    `json:"ScoreThresholdEnabled" name:"ScoreThresholdEnabled"`
		ScoreThreshold        *float64 `json:"ScoreThreshold" name:"ScoreThreshold"`
		Retriever             struct {
			Vector struct {
				TopK                  *int     `json:"TopK" name:"TopK"`
				ScoreThreshold        *float64 `json:"ScoreThreshold" name:"ScoreThreshold"`
				ScoreThresholdEnabled *bool    `json:"ScoreThresholdEnabled" name:"ScoreThresholdEnabled"`
			} `json:"Vector"`
			Inverted struct {
				TopK                  *int     `json:"TopK" name:"TopK"`
				ScoreThreshold        *float64 `json:"ScoreThreshold" name:"ScoreThreshold"`
				ScoreThresholdEnabled *bool    `json:"ScoreThresholdEnabled" name:"ScoreThresholdEnabled"`
			} `json:"Inverted"`
		} `json:"Retriever" name:"Retriever"`
	} `json:"RetrievalModelDict"`
	Tags                  []*string `json:"Tags" name:"Tags"`
	DocForm               *string   `json:"DocForm" name:"DocForm"`
	ExternalKnowledgeInfo struct {
		ExternalKnowledgeId          *string `json:"ExternalKnowledgeId" name:"ExternalKnowledgeId"`
		ExternalKnowledgeApiId       *string `json:"ExternalKnowledgeApiId" name:"ExternalKnowledgeApiId"`
		ExternalKnowledgeApiName     *string `json:"ExternalKnowledgeApiName" name:"ExternalKnowledgeApiName"`
		ExternalKnowledgeApiEndpoint *string `json:"ExternalKnowledgeApiEndpoint" name:"ExternalKnowledgeApiEndpoint"`
	} `json:"ExternalKnowledgeInfo"`
	ExternalRetrievalModel struct {
		TopK                  *int     `json:"TopK" name:"TopK"`
		ScoreThreshold        *float64 `json:"ScoreThreshold" name:"ScoreThreshold"`
		ScoreThresholdEnabled *bool    `json:"ScoreThresholdEnabled" name:"ScoreThresholdEnabled"`
	} `json:"ExternalRetrievalModel"`
	ProjectId   *string `json:"ProjectId" name:"ProjectId"`
	ChargeType  *string `json:"ChargeType" name:"ChargeType"`
	ComputeUnit *int    `json:"ComputeUnit" name:"ComputeUnit"`
}

func (r *DescribeKnowledgeBaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKnowledgeBaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKnowledgeBasesRequest struct {
	*ksyunhttp.BaseRequest
	Page    *int    `json:"Page,omitempty" name:"Page"`
	Limit   *int    `json:"Limit,omitempty" name:"Limit"`
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeKnowledgeBasesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeKnowledgeBasesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Id            *string `json:"Id" name:"Id"`
		Name          *string `json:"Name" name:"Name"`
		DocumentCount *int    `json:"DocumentCount" name:"DocumentCount"`
		CreatedAt     *int64  `json:"CreatedAt" name:"CreatedAt"`
		UpdatedAt     *int64  `json:"UpdatedAt" name:"UpdatedAt"`
		ProjectId     *string `json:"ProjectId" name:"ProjectId"`
		ChargeType    *string `json:"ChargeType" name:"ChargeType"`
		ComputeUnit   *int    `json:"ComputeUnit" name:"ComputeUnit"`
	} `json:"Data"`
	HasMore *bool `json:"HasMore" name:"HasMore"`
	Total   *int  `json:"Total" name:"Total"`
	Page    *int  `json:"Page" name:"Page"`
	Limit   *int  `json:"Limit" name:"Limit"`
}

func (r *DescribeKnowledgeBasesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKnowledgeBasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateKnowledgeBaseRequest struct {
	*ksyunhttp.BaseRequest
	Name              *string                            `json:"Name,omitempty" name:"Name"`
	IndexingTechnique *string                            `json:"IndexingTechnique,omitempty" name:"IndexingTechnique"`
	RetrievalModel    *CreateKnowledgeBaseRetrievalModel `json:"RetrievalModel,omitempty" name:"RetrievalModel"`
	ComputeUnit       *int                               `json:"ComputeUnit,omitempty" name:"ComputeUnit"`
	ProjectId         *string                            `json:"ProjectId,omitempty" name:"ProjectId"`
	ChargeType        *string                            `json:"ChargeType,omitempty" name:"ChargeType"`
}

func (r *CreateKnowledgeBaseRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateKnowledgeBaseResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Id        *string `json:"Id" name:"Id"`
}

func (r *CreateKnowledgeBaseResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateKnowledgeBaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMemorySdkRequest struct {
	*ksyunhttp.BaseRequest
	AgentId            *string              `json:"AgentId,omitempty" name:"AgentId"`
	SessionId          *string              `json:"SessionId,omitempty" name:"SessionId"`
	SceneId            *string              `json:"SceneId,omitempty" name:"SceneId"`
	DataType           *string              `json:"DataType,omitempty" name:"DataType"`
	Data               *CreateMemorySdkData `json:"Data,omitempty" name:"Data"`
	AgentUserId        *string              `json:"AgentUserId,omitempty" name:"AgentUserId"`
	MemoryCollectionId *string              `json:"MemoryCollectionId,omitempty" name:"MemoryCollectionId"`
	Flush              *bool                `json:"Flush,omitempty" name:"Flush"`
}

func (r *CreateMemorySdkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateMemorySdkResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
}

func (r *CreateMemorySdkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMemorySdkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMemorySdkRequest struct {
	*ksyunhttp.BaseRequest
	Query              *string   `json:"Query,omitempty" name:"Query"`
	SceneId            *string   `json:"SceneId,omitempty" name:"SceneId"`
	OccurredAfter      *int64    `json:"OccurredAfter,omitempty" name:"OccurredAfter"`
	OccurredBefore     *int64    `json:"OccurredBefore,omitempty" name:"OccurredBefore"`
	Mode               *string   `json:"Mode,omitempty" name:"Mode"`
	ReturnCitations    *bool     `json:"ReturnCitations,omitempty" name:"ReturnCitations"`
	Limit              *int      `json:"Limit,omitempty" name:"Limit"`
	SceneIds           []*string `json:"SceneIds,omitempty" name:"SceneIds"`
	MemoryCollectionId *string   `json:"MemoryCollectionId,omitempty" name:"MemoryCollectionId"`
	AgentUserId        *string   `json:"AgentUserId,omitempty" name:"AgentUserId"`
}

func (r *QueryMemorySdkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryMemorySdkResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Code      *int    `json:"Code" name:"Code"`
	Message   *string `json:"Message" name:"Message"`
	Data      []struct {
		Memories []struct {
			Query         *string  `json:"Query" name:"Query"`
			TopicId       *string  `json:"TopicId" name:"TopicId"`
			TopicName     *string  `json:"TopicName" name:"TopicName"`
			MemoryId      *string  `json:"MemoryId" name:"MemoryId"`
			Memory        *string  `json:"Memory" name:"Memory"`
			Score         *float64 `json:"Score" name:"Score"`
			OccurredStart *int64   `json:"OccurredStart" name:"OccurredStart"`
			OccurredEnd   *int64   `json:"OccurredEnd" name:"OccurredEnd"`
			Citations     []struct {
				DateType *string `json:"DateType" name:"DateType"`
				AgentId  *string `json:"AgentId" name:"AgentId"`
				Data     struct {
					Datas []struct {
						Text      *string `json:"Text" name:"Text"`
						User      *string `json:"User" name:"User"`
						Timestamp *int64  `json:"Timestamp" name:"Timestamp"`
						MessageId *string `json:"MessageId" name:"MessageId"`
					} `json:"Datas"`
				} `json:"Data" name:"Data"`
			} `json:"Citations"`
		} `json:"Memories" name:"Memories"`
	} `json:"Data"`
}

func (r *QueryMemorySdkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMemorySdkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMemoryCollectionRequest struct {
	*ksyunhttp.BaseRequest
	Name                  *string                                      `json:"Name,omitempty" name:"Name"`
	Description           *string                                      `json:"Description,omitempty" name:"Description"`
	LongTermConfiguration *CreateMemoryCollectionLongTermConfiguration `json:"LongTermConfiguration,omitempty" name:"LongTermConfiguration"`
	MemoryType            *string                                      `json:"MemoryType,omitempty" name:"MemoryType"`
	ProjectId             *string                                      `json:"ProjectId,omitempty" name:"ProjectId"`
	ChargeType            *string                                      `json:"ChargeType,omitempty" name:"ChargeType"`
}

func (r *CreateMemoryCollectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateMemoryCollectionResponse struct {
	*ksyunhttp.BaseResponse
	MemoryCollectionId *string `json:"MemoryCollectionId" name:"MemoryCollectionId"`
	Status             *string `json:"Status" name:"Status"`
}

func (r *CreateMemoryCollectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMemoryCollectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMemoryCollectionRequest struct {
	*ksyunhttp.BaseRequest
	MemoryCollectionId *string `json:"MemoryCollectionId,omitempty" name:"MemoryCollectionId"`
}

func (r *GetMemoryCollectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetMemoryCollectionResponse struct {
	*ksyunhttp.BaseResponse
	MemoryCollectionId    *string `json:"MemoryCollectionId" name:"MemoryCollectionId"`
	Name                  *string `json:"Name" name:"Name"`
	Description           *string `json:"Description" name:"Description"`
	Status                *string `json:"Status" name:"Status"`
	CreateTime            *string `json:"CreateTime" name:"CreateTime"`
	LastUpdateTime        *string `json:"LastUpdateTime" name:"LastUpdateTime"`
	Region                *string `json:"Region" name:"Region"`
	LongTermConfiguration struct {
		Strategies []struct {
			Type *string `json:"Type" name:"Type"`
		} `json:"Strategies" name:"Strategies"`
	} `json:"LongTermConfiguration"`
}

func (r *GetMemoryCollectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMemoryCollectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMemoryCollectionsRequest struct {
	*ksyunhttp.BaseRequest
	CreateTimeAfter    *int64  `json:"CreateTimeAfter,omitempty" name:"CreateTimeAfter"`
	CreateTimeBefore   *int64  `json:"CreateTimeBefore,omitempty" name:"CreateTimeBefore"`
	UpdateTimeAfter    *int64  `json:"UpdateTimeAfter,omitempty" name:"UpdateTimeAfter"`
	UpdateTimeBefore   *int64  `json:"UpdateTimeBefore,omitempty" name:"UpdateTimeBefore"`
	MemoryCollectionId *string `json:"MemoryCollectionId,omitempty" name:"MemoryCollectionId"`
	Name               *string `json:"Name,omitempty" name:"Name"`
	NameKeyword        *string `json:"NameKeyword,omitempty" name:"NameKeyword"`
	Status             *string `json:"Status,omitempty" name:"Status"`
	Marker             *int64  `json:"Marker,omitempty" name:"Marker"`
	MaxResults         *int64  `json:"MaxResults,omitempty" name:"MaxResults"`
}

func (r *ListMemoryCollectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListMemoryCollectionsResponse struct {
	*ksyunhttp.BaseResponse
	Memories []struct {
		MemoryCollectionId    *string `json:"MemoryCollectionId" name:"MemoryCollectionId"`
		Name                  *string `json:"Name" name:"Name"`
		Description           *string `json:"Description" name:"Description"`
		Status                *string `json:"Status" name:"Status"`
		CreateTime            *int64  `json:"CreateTime" name:"CreateTime"`
		LastUpdateTime        *int64  `json:"LastUpdateTime" name:"LastUpdateTime"`
		Region                *string `json:"Region" name:"Region"`
		LongTermConfiguration struct {
			Strategies []struct {
				Type *string `json:"Type" name:"Type"`
			} `json:"Strategies"`
		} `json:"LongTermConfiguration" name:"LongTermConfiguration"`
	} `json:"Memories"`
	MaxResults *int64 `json:"MaxResults" name:"MaxResults"`
	Marker     *int64 `json:"Marker" name:"Marker"`
	Total      *int64 `json:"Total" name:"Total"`
}

func (r *ListMemoryCollectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMemoryCollectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMemoryCollectionRequest struct {
	*ksyunhttp.BaseRequest
	MemoryCollectionId *string `json:"MemoryCollectionId,omitempty" name:"MemoryCollectionId"`
}

func (r *DeleteMemoryCollectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteMemoryCollectionResponse struct {
	*ksyunhttp.BaseResponse
	MemoryCollectionId *string `json:"MemoryCollectionId" name:"MemoryCollectionId"`
}

func (r *DeleteMemoryCollectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMemoryCollectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMemoryBaseServiceRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetMemoryBaseServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetMemoryBaseServiceResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Status    *string `json:"Status" name:"Status"`
}

func (r *GetMemoryBaseServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMemoryBaseServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActivateMemoryBaseServiceRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *ActivateMemoryBaseServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ActivateMemoryBaseServiceResponse struct {
	*ksyunhttp.BaseResponse
	Status    *string `json:"Status" name:"Status"`
	RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ActivateMemoryBaseServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateMemoryBaseServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateMemoryCollectionRequest struct {
	*ksyunhttp.BaseRequest
	MemoryCollectionId    *string                                      `json:"MemoryCollectionId,omitempty" name:"MemoryCollectionId"`
	Description           *string                                      `json:"Description,omitempty" name:"Description"`
	Name                  *string                                      `json:"Name,omitempty" name:"Name"`
	LongTermConfiguration *UpdateMemoryCollectionLongTermConfiguration `json:"LongTermConfiguration,omitempty" name:"LongTermConfiguration"`
}

func (r *UpdateMemoryCollectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type UpdateMemoryCollectionResponse struct {
	*ksyunhttp.BaseResponse
	MemoryCollectionId    *string `json:"MemoryCollectionId" name:"MemoryCollectionId"`
	Name                  *string `json:"Name" name:"Name"`
	LongTermConfiguration struct {
		Strategies []struct {
			Type *string `json:"Type" name:"Type"`
		} `json:"Strategies" name:"Strategies"`
	} `json:"LongTermConfiguration"`
}

func (r *UpdateMemoryCollectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateMemoryCollectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMcpServerRequest struct {
	*ksyunhttp.BaseRequest
	McpServerId *string `json:"McpServerId,omitempty" name:"McpServerId"`
}

func (r *DeleteMcpServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeleteMcpServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *DeleteMcpServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMcpServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMcpServerRequest struct {
	*ksyunhttp.BaseRequest
	McpServerId             *string `json:"McpServerId,omitempty" name:"McpServerId"`
	McpServerName           *string `json:"McpServerName,omitempty" name:"McpServerName"`
	Description             *string `json:"Description,omitempty" name:"Description"`
	Introduction            *string `json:"Introduction,omitempty" name:"Introduction"`
	OutboundAuthFieldValue  *string `json:"OutboundAuthFieldValue,omitempty" name:"OutboundAuthFieldValue"`
	HttpApiConfig           *string `json:"HttpApiConfig,omitempty" name:"HttpApiConfig"`
	HttpApiConfigUpdateType *string `json:"HttpApiConfigUpdateType,omitempty" name:"HttpApiConfigUpdateType"`
}

func (r *ModifyMcpServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyMcpServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *ModifyMcpServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMcpServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMcpServerRequest struct {
	*ksyunhttp.BaseRequest
	McpServerName          *string                          `json:"McpServerName,omitempty" name:"McpServerName"`
	McpServerNameEn        *string                          `json:"McpServerNameEn,omitempty" name:"McpServerNameEn"`
	Description            *string                          `json:"Description,omitempty" name:"Description"`
	Introduction           *string                          `json:"Introduction,omitempty" name:"Introduction"`
	ServiceProtocol        *string                          `json:"ServiceProtocol,omitempty" name:"ServiceProtocol"`
	BackendServiceUrl      *string                          `json:"BackendServiceUrl,omitempty" name:"BackendServiceUrl"`
	AllowCustomAuth        *bool                            `json:"AllowCustomAuth,omitempty" name:"AllowCustomAuth"`
	ServiceCustomHeaders   *string                          `json:"ServiceCustomHeaders,omitempty" name:"ServiceCustomHeaders"`
	OutboundAuthLocation   *string                          `json:"OutboundAuthLocation,omitempty" name:"OutboundAuthLocation"`
	OutboundAuthFieldName  *string                          `json:"OutboundAuthFieldName,omitempty" name:"OutboundAuthFieldName"`
	OutboundAuthFieldValue *string                          `json:"OutboundAuthFieldValue,omitempty" name:"OutboundAuthFieldValue"`
	McpRuntimeConfig       *CreateMcpServerMcpRuntimeConfig `json:"McpRuntimeConfig,omitempty" name:"McpRuntimeConfig"`
	HttpApiConfig          *string                          `json:"HttpApiConfig,omitempty" name:"HttpApiConfig"`
	McpType                *string                          `json:"McpType,omitempty" name:"McpType"`
}

func (r *CreateMcpServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type CreateMcpServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	McpServer struct {
		McpServerId           *string `json:"McpServerId" name:"McpServerId"`
		McpServerName         *string `json:"McpServerName" name:"McpServerName"`
		State                 *string `json:"State" name:"State"`
		CreateTime            *string `json:"CreateTime" name:"CreateTime"`
		McpServerNameEn       *string `json:"McpServerNameEn" name:"McpServerNameEn"`
		Description           *string `json:"Description" name:"Description"`
		Introduction          *string `json:"Introduction" name:"Introduction"`
		ParamConfig           *string `json:"ParamConfig" name:"ParamConfig"`
		ServiceProtocol       *string `json:"ServiceProtocol" name:"ServiceProtocol"`
		IsActivated           *bool   `json:"IsActivated" name:"IsActivated"`
		AllowCustomAuth       *bool   `json:"AllowCustomAuth" name:"AllowCustomAuth"`
		McpType               *string `json:"McpType" name:"McpType"`
		McpStatus             *string `json:"McpStatus" name:"McpStatus"`
		McpStatusMsg          *string `json:"McpStatusMsg" name:"McpStatusMsg"`
		OutboundAuthLocation  *string `json:"OutboundAuthLocation" name:"OutboundAuthLocation"`
		OutboundAuthFieldName *string `json:"OutboundAuthFieldName" name:"OutboundAuthFieldName"`
	} `json:"McpServer"`
}

func (r *CreateMcpServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMcpServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMcpServersRequest struct {
	*ksyunhttp.BaseRequest
	McpServerIds []*string `json:"McpServerIds,omitempty" name:"McpServerIds"`
	NameKeyword  *string   `json:"NameKeyword,omitempty" name:"NameKeyword"`
	Region       *string   `json:"Region,omitempty" name:"Region"`
}

func (r *DescribeMcpServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMcpServersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	McpServers []struct {
		McpServerId           *string `json:"McpServerId" name:"McpServerId"`
		McpServerName         *string `json:"McpServerName" name:"McpServerName"`
		State                 *string `json:"State" name:"State"`
		CreateTime            *string `json:"CreateTime" name:"CreateTime"`
		McpServerNameEn       *string `json:"McpServerNameEn" name:"McpServerNameEn"`
		Description           *string `json:"Description" name:"Description"`
		Introduction          *string `json:"Introduction" name:"Introduction"`
		ParamConfig           *string `json:"ParamConfig" name:"ParamConfig"`
		ServiceProtocol       *string `json:"ServiceProtocol" name:"ServiceProtocol"`
		IsActivated           *bool   `json:"IsActivated" name:"IsActivated"`
		AllowCustomAuth       *bool   `json:"AllowCustomAuth" name:"AllowCustomAuth"`
		McpType               *string `json:"McpType" name:"McpType"`
		McpStatus             *string `json:"McpStatus" name:"McpStatus"`
		McpStatusMsg          *string `json:"McpStatusMsg" name:"McpStatusMsg"`
		OutboundAuthLocation  *string `json:"OutboundAuthLocation" name:"OutboundAuthLocation"`
		OutboundAuthFieldName *string `json:"OutboundAuthFieldName" name:"OutboundAuthFieldName"`
	} `json:"McpServers"`
}

func (r *DescribeMcpServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMcpServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMcpOfficialServersRequest struct {
	*ksyunhttp.BaseRequest
	McpServerIds []*string `json:"McpServerIds,omitempty" name:"McpServerIds"`
	NameKeyword  *string   `json:"NameKeyword,omitempty" name:"NameKeyword"`
	Region       *string   `json:"Region,omitempty" name:"Region"`
}

func (r *DescribeMcpOfficialServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMcpOfficialServersResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	McpServers []struct {
		McpServerId           *string `json:"McpServerId" name:"McpServerId"`
		McpServerName         *string `json:"McpServerName" name:"McpServerName"`
		State                 *string `json:"State" name:"State"`
		CreateTime            *string `json:"CreateTime" name:"CreateTime"`
		McpServerNameEn       *string `json:"McpServerNameEn" name:"McpServerNameEn"`
		Description           *string `json:"Description" name:"Description"`
		Introduction          *string `json:"Introduction" name:"Introduction"`
		ParamConfig           *string `json:"ParamConfig" name:"ParamConfig"`
		ServiceProtocol       *string `json:"ServiceProtocol" name:"ServiceProtocol"`
		IsActivated           *bool   `json:"IsActivated" name:"IsActivated"`
		AllowCustomAuth       *bool   `json:"AllowCustomAuth" name:"AllowCustomAuth"`
		McpType               *string `json:"McpType" name:"McpType"`
		McpStatus             *string `json:"McpStatus" name:"McpStatus"`
		McpStatusMsg          *string `json:"McpStatusMsg" name:"McpStatusMsg"`
		OutboundAuthLocation  *string `json:"OutboundAuthLocation" name:"OutboundAuthLocation"`
		OutboundAuthFieldName *string `json:"OutboundAuthFieldName" name:"OutboundAuthFieldName"`
	} `json:"McpServers"`
}

func (r *DescribeMcpOfficialServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMcpOfficialServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeactivateMcpOfficialServerRequest struct {
	*ksyunhttp.BaseRequest
	McpServerId *string `json:"McpServerId,omitempty" name:"McpServerId"`
}

func (r *DeactivateMcpOfficialServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DeactivateMcpOfficialServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *DeactivateMcpOfficialServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeactivateMcpOfficialServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActivateMcpOfficialServerRequest struct {
	*ksyunhttp.BaseRequest
	McpServerId    *string `json:"McpServerId,omitempty" name:"McpServerId"`
	AuthFieldValue *string `json:"AuthFieldValue,omitempty" name:"AuthFieldValue"`
}

func (r *ActivateMcpOfficialServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ActivateMcpOfficialServerResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *ActivateMcpOfficialServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateMcpOfficialServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMcpSquaresRequest struct {
	*ksyunhttp.BaseRequest
	McpServerIds []*string `json:"McpServerIds,omitempty" name:"McpServerIds"`
	NameKeyword  *string   `json:"NameKeyword,omitempty" name:"NameKeyword"`
}

func (r *DescribeMcpSquaresRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMcpSquaresResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	McpServers []struct {
		McpServerId           *string `json:"McpServerId" name:"McpServerId"`
		McpServerName         *string `json:"McpServerName" name:"McpServerName"`
		State                 *string `json:"State" name:"State"`
		CreateTime            *string `json:"CreateTime" name:"CreateTime"`
		McpServerNameEn       *string `json:"McpServerNameEn" name:"McpServerNameEn"`
		Description           *string `json:"Description" name:"Description"`
		Introduction          *string `json:"Introduction" name:"Introduction"`
		ParamConfig           *string `json:"ParamConfig" name:"ParamConfig"`
		ServiceProtocol       *string `json:"ServiceProtocol" name:"ServiceProtocol"`
		IsActivated           *bool   `json:"IsActivated" name:"IsActivated"`
		AllowCustomAuth       *bool   `json:"AllowCustomAuth" name:"AllowCustomAuth"`
		McpType               *string `json:"McpType" name:"McpType"`
		McpStatus             *string `json:"McpStatus" name:"McpStatus"`
		McpStatusMsg          *string `json:"McpStatusMsg" name:"McpStatusMsg"`
		OutboundAuthLocation  *string `json:"OutboundAuthLocation" name:"OutboundAuthLocation"`
		OutboundAuthFieldName *string `json:"OutboundAuthFieldName" name:"OutboundAuthFieldName"`
	} `json:"McpServers"`
}

func (r *DescribeMcpSquaresResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMcpSquaresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMcpOfficialServerDetailRequest struct {
	*ksyunhttp.BaseRequest
	McpServerId *string `json:"McpServerId,omitempty" name:"McpServerId"`
}

func (r *GetMcpOfficialServerDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetMcpOfficialServerDetailResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	McpServer struct {
		McpServerId           *string `json:"McpServerId" name:"McpServerId"`
		McpServerName         *string `json:"McpServerName" name:"McpServerName"`
		State                 *string `json:"State" name:"State"`
		CreateTime            *string `json:"CreateTime" name:"CreateTime"`
		McpServerNameEn       *string `json:"McpServerNameEn" name:"McpServerNameEn"`
		Description           *string `json:"Description" name:"Description"`
		Introduction          *string `json:"Introduction" name:"Introduction"`
		ParamConfig           *string `json:"ParamConfig" name:"ParamConfig"`
		ServiceProtocol       *string `json:"ServiceProtocol" name:"ServiceProtocol"`
		Tools                 *string `json:"Tools" name:"Tools"`
		IsActivated           *bool   `json:"IsActivated" name:"IsActivated"`
		AllowCustomAuth       *bool   `json:"AllowCustomAuth" name:"AllowCustomAuth"`
		McpType               *string `json:"McpType" name:"McpType"`
		McpStatus             *string `json:"McpStatus" name:"McpStatus"`
		McpStatusMsg          *string `json:"McpStatusMsg" name:"McpStatusMsg"`
		OutboundAuthLocation  *string `json:"OutboundAuthLocation" name:"OutboundAuthLocation"`
		OutboundAuthFieldName *string `json:"OutboundAuthFieldName" name:"OutboundAuthFieldName"`
	} `json:"McpServer"`
}

func (r *GetMcpOfficialServerDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMcpOfficialServerDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMcpServerDetailRequest struct {
	*ksyunhttp.BaseRequest
	McpServerId *string `json:"McpServerId,omitempty" name:"McpServerId"`
}

func (r *GetMcpServerDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetMcpServerDetailResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	McpServer struct {
		McpServerId           *string `json:"McpServerId" name:"McpServerId"`
		McpServerName         *string `json:"McpServerName" name:"McpServerName"`
		State                 *string `json:"State" name:"State"`
		CreateTime            *string `json:"CreateTime" name:"CreateTime"`
		McpServerNameEn       *string `json:"McpServerNameEn" name:"McpServerNameEn"`
		Description           *string `json:"Description" name:"Description"`
		Introduction          *string `json:"Introduction" name:"Introduction"`
		ParamConfig           *string `json:"ParamConfig" name:"ParamConfig"`
		ServiceProtocol       *string `json:"ServiceProtocol" name:"ServiceProtocol"`
		Tools                 *string `json:"Tools" name:"Tools"`
		IsActivated           *bool   `json:"IsActivated" name:"IsActivated"`
		AllowCustomAuth       *bool   `json:"AllowCustomAuth" name:"AllowCustomAuth"`
		McpType               *string `json:"McpType" name:"McpType"`
		McpStatus             *string `json:"McpStatus" name:"McpStatus"`
		McpStatusMsg          *string `json:"McpStatusMsg" name:"McpStatusMsg"`
		OutboundAuthLocation  *string `json:"OutboundAuthLocation" name:"OutboundAuthLocation"`
		OutboundAuthFieldName *string `json:"OutboundAuthFieldName" name:"OutboundAuthFieldName"`
	} `json:"McpServer"`
}

func (r *GetMcpServerDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMcpServerDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMcpSquareDetailRequest struct {
	*ksyunhttp.BaseRequest
	McpServerId *string `json:"McpServerId,omitempty" name:"McpServerId"`
}

func (r *GetMcpSquareDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetMcpSquareDetailResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	McpServer struct {
		McpServerId           *string `json:"McpServerId" name:"McpServerId"`
		McpServerName         *string `json:"McpServerName" name:"McpServerName"`
		State                 *string `json:"State" name:"State"`
		CreateTime            *string `json:"CreateTime" name:"CreateTime"`
		McpServerNameEn       *string `json:"McpServerNameEn" name:"McpServerNameEn"`
		Description           *string `json:"Description" name:"Description"`
		Introduction          *string `json:"Introduction" name:"Introduction"`
		ParamConfig           *string `json:"ParamConfig" name:"ParamConfig"`
		ServiceProtocol       *string `json:"ServiceProtocol" name:"ServiceProtocol"`
		Tools                 *string `json:"Tools" name:"Tools"`
		IsActivated           *bool   `json:"IsActivated" name:"IsActivated"`
		AllowCustomAuth       *bool   `json:"AllowCustomAuth" name:"AllowCustomAuth"`
		McpType               *string `json:"McpType" name:"McpType"`
		McpStatus             *string `json:"McpStatus" name:"McpStatus"`
		McpStatusMsg          *string `json:"McpStatusMsg" name:"McpStatusMsg"`
		OutboundAuthLocation  *string `json:"OutboundAuthLocation" name:"OutboundAuthLocation"`
		OutboundAuthFieldName *string `json:"OutboundAuthFieldName" name:"OutboundAuthFieldName"`
	} `json:"McpServer"`
}

func (r *GetMcpSquareDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMcpSquareDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListSessionsRequest struct {
	*ksyunhttp.BaseRequest
	MemoryCollectionId *string `json:"MemoryCollectionId,omitempty" name:"MemoryCollectionId"`
	AgentUserId        *string `json:"AgentUserId,omitempty" name:"AgentUserId"`
	Query              *string `json:"Query,omitempty" name:"Query"`
	Page               *int    `json:"Page,omitempty" name:"Page"`
	PageSize           *int    `json:"PageSize,omitempty" name:"PageSize"`
	CreatedAfter       *int    `json:"CreatedAfter,omitempty" name:"CreatedAfter"`
	CreatedBefore      *int    `json:"CreatedBefore,omitempty" name:"CreatedBefore"`
}

func (r *ListSessionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListSessionsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      struct {
		Total *int `json:"Total" name:"Total"`
		Items []struct {
			DataId *string `json:"DataId" name:"DataId"`
			Data   struct {
				Conversation []struct {
					Role      *string `json:"Role" name:"Role"`
					Text      *string `json:"Text" name:"Text"`
					CreatedAt *int    `json:"CreatedAt" name:"CreatedAt"`
				} `json:"Conversation" name:"Conversation"`
				State *int `json:"State" name:"State"`
			} `json:"Data"`
			CreatedAt *int `json:"CreatedAt" name:"CreatedAt"`
		} `json:"Items" name:"Items"`
	} `json:"Data"`
}

func (r *ListSessionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddSessionRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *AddSessionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type AddSessionResponse struct {
	*ksyunhttp.BaseResponse
	Code    *int    `json:"Code" name:"Code"`
	Message *string `json:"Message" name:"Message"`
	Data    *string `json:"Data" name:"Data"`
}

func (r *AddSessionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMemoryCollectionMetricsRequest struct {
	*ksyunhttp.BaseRequest
	MemoryCollectionId *string `json:"MemoryCollectionId,omitempty" name:"MemoryCollectionId"`
	StartTime          *int64  `json:"StartTime,omitempty" name:"StartTime"`
	EndTime            *int64  `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *QueryMemoryCollectionMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryMemoryCollectionMetricsResponse struct {
	*ksyunhttp.BaseResponse
	MemoryCollectionId *string `json:"MemoryCollectionId" name:"MemoryCollectionId"`
	MetricsByAction    []struct {
		QueryMemory struct {
			RequestCount *int     `json:"RequestCount" name:"RequestCount"`
			AvgDuration  *float64 `json:"AvgDuration" name:"AvgDuration"`
			ErrorRate    *float64 `json:"ErrorRate" name:"ErrorRate"`
		} `json:"QueryMemory" name:"QueryMemory"`
		CreateMemory struct {
			RequestCount *int     `json:"RequestCount" name:"RequestCount"`
			AvgDuration  *int     `json:"AvgDuration" name:"AvgDuration"`
			ErrorRate    *float64 `json:"ErrorRate" name:"ErrorRate"`
		} `json:"CreateMemory" name:"CreateMemory"`
	} `json:"MetricsByAction"`
}

func (r *QueryMemoryCollectionMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMemoryCollectionMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuerySessionMemoriesRequest struct {
	*ksyunhttp.BaseRequest
	MemoryCollectionId *string `json:"MemoryCollectionId,omitempty" name:"MemoryCollectionId"`
	SessionId          *string `json:"SessionId,omitempty" name:"SessionId"`
}

func (r *QuerySessionMemoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QuerySessionMemoriesResponse struct {
	*ksyunhttp.BaseResponse
	Data []struct {
		TopicId       *string `json:"TopicId" name:"TopicId"`
		TopicName     *string `json:"TopicName" name:"TopicName"`
		MemoryId      *string `json:"MemoryId" name:"MemoryId"`
		Memory        *string `json:"Memory" name:"Memory"`
		OccurredStart *int    `json:"OccurredStart" name:"OccurredStart"`
		OccurredEnd   *int    `json:"OccurredEnd" name:"OccurredEnd"`
	} `json:"Data"`
	Code    *int    `json:"Code" name:"Code"`
	Message *string `json:"Message" name:"Message"`
}

func (r *QuerySessionMemoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QuerySessionMemoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetrieveHistoriesRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId *string `json:"DatasetId,omitempty" name:"DatasetId"`
	Page      *int    `json:"Page,omitempty" name:"Page"`
	Limit     *int    `json:"Limit,omitempty" name:"Limit"`
}

func (r *RetrieveHistoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type RetrieveHistoriesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Data      []struct {
		Id        *string `json:"Id" name:"Id"`
		Content   *string `json:"Content" name:"Content"`
		CreatedAt *int64  `json:"CreatedAt" name:"CreatedAt"`
	} `json:"Data"`
	HasMore *bool `json:"HasMore" name:"HasMore"`
	Total   *int  `json:"Total" name:"Total"`
	Page    *int  `json:"Page" name:"Page"`
	Limit   *int  `json:"Limit" name:"Limit"`
}

func (r *RetrieveHistoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetrieveHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReindexDocumentsRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId   *string   `json:"DatasetId,omitempty" name:"DatasetId"`
	DocumentIds []*string `json:"DocumentIds,omitempty" name:"DocumentIds"`
}

func (r *ReindexDocumentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ReindexDocumentsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *ReindexDocumentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReindexDocumentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDocumentStatusRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId  *string `json:"DatasetId,omitempty" name:"DatasetId"`
	DocumentId *string `json:"DocumentId,omitempty" name:"DocumentId"`
	Status     *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyDocumentStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ModifyDocumentStatusResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Result    *bool   `json:"Result" name:"Result"`
}

func (r *ModifyDocumentStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDocumentStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApiDetailRequest struct {
	*ksyunhttp.BaseRequest
	ApiService *string `json:"ApiService,omitempty" name:"ApiService"`
	ApiName    *string `json:"ApiName,omitempty" name:"ApiName"`
	ApiVersion *string `json:"ApiVersion,omitempty" name:"ApiVersion"`
}

func (r *GetApiDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetApiDetailResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	ApiDetail struct {
		ApiService      *string `json:"ApiService" name:"ApiService"`
		ApiName         *string `json:"ApiName" name:"ApiName"`
		ApiVersion      *string `json:"ApiVersion" name:"ApiVersion"`
		ApiNameCn       *string `json:"ApiNameCn" name:"ApiNameCn"`
		ApiDescription  *string `json:"ApiDescription" name:"ApiDescription"`
		ApiInstructions *string `json:"ApiInstructions" name:"ApiInstructions"`
		ApiParams       struct {
			Type        *string `json:"type" name:"type"`
			Description *string `json:"description" name:"description"`
			Properties  struct {
				Type        *string `json:"Type" name:"Type"`
				Format      *string `json:"Format" name:"Format"`
				Description *string `json:"Description" name:"Description"`
				MaxLength   *int    `json:"MaxLength" name:"MaxLength"`
				MinLength   *int    `json:"MinLength" name:"MinLength"`
				MinItems    *int    `json:"MinItems" name:"MinItems"`
				MaxItems    *int    `json:"MaxItems" name:"MaxItems"`
				Properties  struct {
				} `json:"Properties" name:"Properties"`
				Items struct {
				} `json:"Items" name:"Items"`
				Required             []*string `json:"Required" name:"Required"`
				AdditionalProperties *bool     `json:"AdditionalProperties" name:"AdditionalProperties"`
			} `json:"Properties"`
			Required             []*string `json:"Required" name:"Required"`
			AdditionalProperties *bool     `json:"additionalProperties" name:"additionalProperties"`
		} `json:"ApiParams" name:"ApiParams"`
		HttpMethod     *string `json:"HttpMethod" name:"HttpMethod"`
		ContentType    *string `json:"ContentType" name:"ContentType"`
		ApiGroup       *string `json:"ApiGroup" name:"ApiGroup"`
		ApiExplorerUrl *string `json:"ApiExplorerUrl" name:"ApiExplorerUrl"`
		OperationType  *string `json:"OperationType" name:"OperationType"`
	} `json:"ApiDetail"`
}

func (r *GetApiDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApiDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetApiOverviewRequest struct {
	*ksyunhttp.BaseRequest
	ApiService *string `json:"ApiService,omitempty" name:"ApiService"`
	ApiVersion *string `json:"ApiVersion,omitempty" name:"ApiVersion"`
}

func (r *GetApiOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetApiOverviewResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	ApiOverviewList []struct {
		ApiService      *string `json:"ApiService" name:"ApiService"`
		ApiName         *string `json:"ApiName" name:"ApiName"`
		ApiVersion      *string `json:"ApiVersion" name:"ApiVersion"`
		ApiNameCn       *string `json:"ApiNameCn" name:"ApiNameCn"`
		ApiDescription  *string `json:"ApiDescription" name:"ApiDescription"`
		ApiInstructions *string `json:"ApiInstructions" name:"ApiInstructions"`
		ApiGroup        *string `json:"ApiGroup" name:"ApiGroup"`
		ApiExplorerUrl  *string `json:"ApiExplorerUrl" name:"ApiExplorerUrl"`
		OperationType   *string `json:"OperationType" name:"OperationType"`
	} `json:"ApiOverviewList"`
}

func (r *GetApiOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetApiOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetProductListRequest struct {
	*ksyunhttp.BaseRequest
}

func (r *GetProductListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type GetProductListResponse struct {
	*ksyunhttp.BaseResponse
	RequestId       *string `json:"RequestId" name:"RequestId"`
	ProductInfoList []struct {
		ApiService     *string   `json:"ApiService" name:"ApiService"`
		ApiServiceName *string   `json:"ApiServiceName" name:"ApiServiceName"`
		Versions       []*string `json:"Versions" name:"Versions"`
	} `json:"ProductInfoList"`
}

func (r *GetProductListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetProductListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMcpRuntimeMetricsRequest struct {
	*ksyunhttp.BaseRequest
	McpServerId *string `json:"McpServerId,omitempty" name:"McpServerId"`
	StartTime   *int64  `json:"StartTime,omitempty" name:"StartTime"`
	EndTime     *int64  `json:"EndTime,omitempty" name:"EndTime"`
	Interval    *int    `json:"Interval,omitempty" name:"Interval"`
}

func (r *DescribeMcpRuntimeMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMcpRuntimeMetricsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId     *string `json:"RequestId" name:"RequestId"`
	McpServerId   *string `json:"McpServerId" name:"McpServerId"`
	StartTime     *int64  `json:"StartTime" name:"StartTime"`
	EndTime       *int64  `json:"EndTime" name:"EndTime"`
	Interval      *int    `json:"Interval" name:"Interval"`
	InstanceCount *int    `json:"InstanceCount" name:"InstanceCount"`
	Series        []struct {
		MetricName  *string `json:"MetricName" name:"MetricName"`
		DisplayName *string `json:"DisplayName" name:"DisplayName"`
		Unit        *string `json:"Unit" name:"Unit"`
		Points      []struct {
			Timestamp *int64   `json:"Timestamp" name:"Timestamp"`
			Value     *float64 `json:"Value" name:"Value"`
		} `json:"Points" name:"Points"`
	} `json:"Series"`
}

func (r *DescribeMcpRuntimeMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMcpRuntimeMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMcpMetricsRequest struct {
	*ksyunhttp.BaseRequest
	StartTime   *int64  `json:"StartTime,omitempty" name:"StartTime"`
	EndTime     *int64  `json:"EndTime,omitempty" name:"EndTime"`
	Interval    *int    `json:"Interval,omitempty" name:"Interval"`
	McpType     *string `json:"McpType,omitempty" name:"McpType"`
	McpServerId *string `json:"McpServerId,omitempty" name:"McpServerId"`
}

func (r *QueryMcpMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type QueryMcpMetricsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId *string `json:"RequestId" name:"RequestId"`
	Metrics   []struct {
		McpServerId *string `json:"McpServerId" name:"McpServerId"`
		Interval    *int    `json:"Interval" name:"Interval"`
		StartTime   *int64  `json:"StartTime" name:"StartTime"`
		EndTime     *int64  `json:"EndTime" name:"EndTime"`
		Data        []struct {
			MetricName  *string `json:"MetricName" name:"MetricName"`
			DisplayName *string `json:"DisplayName" name:"DisplayName"`
			Unit        *string `json:"Unit" name:"Unit"`
			Points      []struct {
				Timestamp *int64   `json:"Timestamp" name:"Timestamp"`
				Value     *float64 `json:"Value" name:"Value"`
			} `json:"Points"`
			StatusPoint struct {
			} `json:"StatusPoint"`
		} `json:"Data" name:"Data"`
	} `json:"Metrics"`
}

func (r *QueryMcpMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMcpMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKnowledgeTokenMonitorRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId   *string `json:"DatasetId,omitempty" name:"DatasetId"`
	StartTime   *int64  `json:"StartTime,omitempty" name:"StartTime"`
	EndTime     *int64  `json:"EndTime,omitempty" name:"EndTime"`
	Granularity *string `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *DescribeKnowledgeTokenMonitorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeKnowledgeTokenMonitorResponse struct {
	*ksyunhttp.BaseResponse
	RequestId                *string `json:"RequestId" name:"RequestId"`
	EmbeddingInputTokens     *int64  `json:"EmbeddingInputTokens" name:"EmbeddingInputTokens"`
	RerankInputTokens        *int64  `json:"RerankInputTokens" name:"RerankInputTokens"`
	EmbeddingInputTokenTrend []struct {
		Timestamp *int64 `json:"Timestamp" name:"Timestamp"`
		Value     *int64 `json:"Value" name:"Value"`
	} `json:"EmbeddingInputTokenTrend"`
	RerankInputTokenTrend []struct {
		Timestamp *int64 `json:"Timestamp" name:"Timestamp"`
		Value     *int64 `json:"Value" name:"Value"`
	} `json:"RerankInputTokenTrend"`
	RequestCountTrend []struct {
		Timestamp *int64 `json:"Timestamp" name:"Timestamp"`
		Value     *int64 `json:"Value" name:"Value"`
	} `json:"RequestCountTrend"`
	RpmRateLimits []struct {
		Model *string  `json:"Model" name:"Model"`
		Value *float64 `json:"Value" name:"Value"`
	} `json:"RpmRateLimits"`
	TpmRateLimits []struct {
		Model *string  `json:"Model" name:"Model"`
		Value *float64 `json:"Value" name:"Value"`
	} `json:"TpmRateLimits"`
}

func (r *DescribeKnowledgeTokenMonitorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKnowledgeTokenMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKnowledgeStorageMonitorRequest struct {
	*ksyunhttp.BaseRequest
	DatasetId   *string `json:"DatasetId,omitempty" name:"DatasetId"`
	StartTime   *int64  `json:"StartTime,omitempty" name:"StartTime"`
	EndTime     *int64  `json:"EndTime,omitempty" name:"EndTime"`
	Granularity *string `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *DescribeKnowledgeStorageMonitorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeKnowledgeStorageMonitorResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string  `json:"RequestId" name:"RequestId"`
	CurrentStorage *float64 `json:"CurrentStorage" name:"CurrentStorage"`
	TodayIncrement *float64 `json:"TodayIncrement" name:"TodayIncrement"`
	DocumentCount  *int     `json:"DocumentCount" name:"DocumentCount"`
	StorageTrend   []struct {
		Timestamp *int64   `json:"Timestamp" name:"Timestamp"`
		Value     *float64 `json:"Value" name:"Value"`
	} `json:"StorageTrend"`
}

func (r *DescribeKnowledgeStorageMonitorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKnowledgeStorageMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMcpRuntimeLogsRequest struct {
	*ksyunhttp.BaseRequest
	McpServerId *string `json:"McpServerId,omitempty" name:"McpServerId"`
	StartTime   *int64  `json:"StartTime,omitempty" name:"StartTime"`
	EndTime     *int64  `json:"EndTime,omitempty" name:"EndTime"`
	Keyword     *string `json:"Keyword,omitempty" name:"Keyword"`
	Page        *int    `json:"Page,omitempty" name:"Page"`
	Limit       *int    `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMcpRuntimeLogsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMcpRuntimeLogsResponse struct {
	*ksyunhttp.BaseResponse
	RequestId   *string   `json:"RequestId" name:"RequestId"`
	McpServerId *string   `json:"McpServerId" name:"McpServerId"`
	StartTime   *int64    `json:"StartTime" name:"StartTime"`
	EndTime     *int64    `json:"EndTime" name:"EndTime"`
	LogType     *string   `json:"LogType" name:"LogType"`
	Logs        []*string `json:"Logs" name:"Logs"`
	Total       *int64    `json:"Total" name:"Total"`
	Page        *int      `json:"Page" name:"Page"`
	Limit       *int      `json:"Limit" name:"Limit"`
}

func (r *DescribeMcpRuntimeLogsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMcpRuntimeLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMemoryTokenMonitorRequest struct {
	*ksyunhttp.BaseRequest
	MemoryId    *string `json:"MemoryId,omitempty" name:"MemoryId"`
	StartTime   *int64  `json:"StartTime,omitempty" name:"StartTime"`
	EndTime     *int64  `json:"EndTime,omitempty" name:"EndTime"`
	Granularity *string `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *DescribeMemoryTokenMonitorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMemoryTokenMonitorResponse struct {
	*ksyunhttp.BaseResponse
	RequestId            *string `json:"RequestId" name:"RequestId"`
	EmbeddingInputTokens *int64  `json:"EmbeddingInputTokens" name:"EmbeddingInputTokens"`
	RerankInputTokens    *int64  `json:"RerankInputTokens" name:"RerankInputTokens"`
	ChatOutputTokens     *int64  `json:"ChatOutputTokens" name:"ChatOutputTokens"`
	ChatHitCacheTokens   *int64  `json:"ChatHitCacheTokens" name:"ChatHitCacheTokens"`
	ChatMissCacheTokens  *int64  `json:"ChatMissCacheTokens" name:"ChatMissCacheTokens"`
	RequestCountTrend    []struct {
		Timestamp *int64 `json:"Timestamp" name:"Timestamp"`
		Value     *int64 `json:"Value" name:"Value"`
	} `json:"RequestCountTrend"`
	EmbeddingInputTokenTrend []struct {
		Timestamp *int64 `json:"Timestamp" name:"Timestamp"`
		Value     *int64 `json:"Value" name:"Value"`
	} `json:"EmbeddingInputTokenTrend"`
	RerankInputTokenTrend []struct {
		Timestamp *int64 `json:"Timestamp" name:"Timestamp"`
		Value     *int64 `json:"Value" name:"Value"`
	} `json:"RerankInputTokenTrend"`
	ChatOutputCountTrend []struct {
		Timestamp *int64 `json:"Timestamp" name:"Timestamp"`
		Value     *int64 `json:"Value" name:"Value"`
	} `json:"ChatOutputCountTrend"`
	ChatHitCacheCountTrend []struct {
		Timestamp *int64 `json:"Timestamp" name:"Timestamp"`
		Value     *int64 `json:"Value" name:"Value"`
	} `json:"ChatHitCacheCountTrend"`
	ChatMissCacheCountTrend []struct {
		Timestamp *int64 `json:"Timestamp" name:"Timestamp"`
		Value     *int64 `json:"Value" name:"Value"`
	} `json:"ChatMissCacheCountTrend"`
	RpmRateLimits []struct {
		Model *string  `json:"Model" name:"Model"`
		Value *float64 `json:"Value" name:"Value"`
	} `json:"RpmRateLimits"`
	TpmRateLimits []struct {
		Model *string  `json:"Model" name:"Model"`
		Value *float64 `json:"Value" name:"Value"`
	} `json:"TpmRateLimits"`
}

func (r *DescribeMemoryTokenMonitorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMemoryTokenMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMemoryStorageMonitorRequest struct {
	*ksyunhttp.BaseRequest
	MemoryId    *string `json:"MemoryId,omitempty" name:"MemoryId"`
	StartTime   *int64  `json:"StartTime,omitempty" name:"StartTime"`
	EndTime     *int64  `json:"EndTime,omitempty" name:"EndTime"`
	Granularity *string `json:"Granularity,omitempty" name:"Granularity"`
}

func (r *DescribeMemoryStorageMonitorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type DescribeMemoryStorageMonitorResponse struct {
	*ksyunhttp.BaseResponse
	RequestId      *string  `json:"RequestId" name:"RequestId"`
	CurrentStorage *float64 `json:"CurrentStorage" name:"CurrentStorage"`
	TodayIncrement *float64 `json:"TodayIncrement" name:"TodayIncrement"`
	StorageTrend   []struct {
		Timestamp *int64   `json:"Timestamp" name:"Timestamp"`
		Value     *float64 `json:"Value" name:"Value"`
	} `json:"StorageTrend"`
}

func (r *DescribeMemoryStorageMonitorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMemoryStorageMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListMemoriesRequest struct {
	*ksyunhttp.BaseRequest
	MemoryCollectionId *string `json:"MemoryCollectionId,omitempty" name:"MemoryCollectionId"`
	AgentUserId        *string `json:"AgentUserId,omitempty" name:"AgentUserId"`
	TopicId            *string `json:"TopicId,omitempty" name:"TopicId"`
	Query              *string `json:"Query,omitempty" name:"Query"`
	Page               *int64  `json:"Page,omitempty" name:"Page"`
	PageSize           *int64  `json:"PageSize,omitempty" name:"PageSize"`
	SortBy             *string `json:"SortBy,omitempty" name:"SortBy"`
	SortOrder          *string `json:"SortOrder,omitempty" name:"SortOrder"`
	CreatedAfter       *int64  `json:"CreatedAfter,omitempty" name:"CreatedAfter"`
	CreatedBefore      *int64  `json:"CreatedBefore,omitempty" name:"CreatedBefore"`
	OccurredAfter      *int64  `json:"OccurredAfter,omitempty" name:"OccurredAfter"`
	OccurredBefore     *int64  `json:"OccurredBefore,omitempty" name:"OccurredBefore"`
}

func (r *ListMemoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListMemoriesResponse struct {
	*ksyunhttp.BaseResponse
	RequestId  *string `json:"RequestId" name:"RequestId"`
	Total      *int64  `json:"Total" name:"Total"`
	MemoryList []struct {
		MemoryId           *string `json:"MemoryId" name:"MemoryId"`
		Memory             *string `json:"Memory" name:"Memory"`
		SourceType         *string `json:"SourceType" name:"SourceType"`
		TopicId            *string `json:"TopicId" name:"TopicId"`
		TopicName          *string `json:"TopicName" name:"TopicName"`
		MemoryCollectionId *string `json:"MemoryCollectionId" name:"MemoryCollectionId"`
		AgentUserId        *string `json:"AgentUserId" name:"AgentUserId"`
		CreatedAt          *int64  `json:"CreatedAt" name:"CreatedAt"`
		UpdatedAt          *int64  `json:"UpdatedAt" name:"UpdatedAt"`
		OccurredStart      *int64  `json:"OccurredStart" name:"OccurredStart"`
		OccurredEnd        *int64  `json:"OccurredEnd" name:"OccurredEnd"`
	} `json:"MemoryList"`
}

func (r *ListMemoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListMemoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
