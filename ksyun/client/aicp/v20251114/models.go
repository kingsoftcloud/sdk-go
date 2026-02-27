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
	WordCount              *int    `json:"WordCount" name:"WordCount"`
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
		WordCount     *int    `json:"WordCount" name:"WordCount"`
		CreatedAt     *int64  `json:"CreatedAt" name:"CreatedAt"`
		UpdatedAt     *int64  `json:"UpdatedAt" name:"UpdatedAt"`
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
