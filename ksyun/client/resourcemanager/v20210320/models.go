package v20210320
import (
	"encoding/json"
    "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
    ksyunhttp "github.com/kingsoftcloud/sdk-go/ksyun/common/http"
)

type CreateFolderRequest struct {
    *ksyunhttp.BaseRequest
    ParentId *string `json:"ParentId,omitempty" name:"ParentId"`
    Name *string `json:"Name,omitempty" name:"Name"`
    Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFolderRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "CreateFolderRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type CreateFolderResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *CreateFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFolderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFolderRequest struct {
    *ksyunhttp.BaseRequest
    FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

func (r *DeleteFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFolderRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "DeleteFolderRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFolderResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *DeleteFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFolderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFolderRequest struct {
    *ksyunhttp.BaseRequest
    FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
    ParentId *string `json:"ParentId,omitempty" name:"ParentId"`
    Name *string `json:"Name,omitempty" name:"Name"`
    Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *UpdateFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFolderRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateFolderRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFolderResponse struct {
    *ksyunhttp.BaseResponse
    Result *string `json:"Result" name:"Result"`
}

func (r *UpdateFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFolderResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAccountsForParentRequest struct {
    *ksyunhttp.BaseRequest
    FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
    Search *string `json:"Search,omitempty" name:"Search"`
    Page *int `json:"Page,omitempty" name:"Page"`
    PageSize *int `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListAccountsForParentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAccountsForParentRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListAccountsForParentRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListAccountsForParentResponse struct {
    *ksyunhttp.BaseResponse
	Members []struct {
		CreatedTime *string `json:"CreatedTime"`
		Name *string `json:"Name"`
		AdminPermission *int `json:"AdminPermission"`
		UserId *int `json:"UserId"`
		UserName *string `json:"UserName"`
		UserType *int `json:"UserType"`
		Phone *string `json:"Phone"`
		FolderId *string `json:"FolderId"`
	} `json:"Members"`
    Count *int `json:"Count" name:"Count"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListAccountsForParentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAccountsForParentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MoveAccountRequest struct {
    *ksyunhttp.BaseRequest
    Ids *string `json:"Ids,omitempty" name:"Ids"`
    FromFolderId *string `json:"FromFolderId,omitempty" name:"FromFolderId"`
    ToFolderId *string `json:"ToFolderId,omitempty" name:"ToFolderId"`
}

func (r *MoveAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MoveAccountRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "MoveAccountRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type MoveAccountResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *MoveAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *MoveAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateAccountRequest struct {
    *ksyunhttp.BaseRequest
    MemberId *int `json:"MemberId,omitempty" name:"MemberId"`
    NewDisplayName *string `json:"NewDisplayName,omitempty" name:"NewDisplayName"`
    FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

func (r *UpdateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAccountRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "UpdateAccountRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type UpdateAccountResponse struct {
    *ksyunhttp.BaseResponse
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *UpdateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAccountResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAccountsRequest struct {
    *ksyunhttp.BaseRequest
    PageNumber *int `json:"PageNumber,omitempty" name:"PageNumber"`
    PageSize *int `json:"PageSize,omitempty" name:"PageSize"`
    IsAll *int `json:"IsAll,omitempty" name:"IsAll"`
}

func (r *ListAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAccountsRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListAccountsRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListAccountsResponse struct {
    *ksyunhttp.BaseResponse
	Accounts []struct {
		AccountId *int `json:"AccountId"`
		AccountName *string `json:"AccountName"`
		DisplayName *string `json:"DisplayName"`
		Type *string `json:"Type"`
		JoinedTime *string `json:"JoinedTime"`
	} `json:"Accounts"`
	Pagination struct {
		TotalCount *int `json:"TotalCount"`
		PageNumber *string `json:"PageNumber"`
		PageSize *string `json:"PageSize"`
	} `json:"Pagination"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAccountsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListFoldersRequest struct {
    *ksyunhttp.BaseRequest
}

func (r *ListFoldersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListFoldersRequest) FromJsonString(s string) error {
    f := make(map[string]interface{})
    if err := json.Unmarshal([]byte(s), &f); err != nil {
        return err
    }
    if len(f) > 0 {
        return errors.NewKsyunSDKError("ClientError.BuildRequestError", "ListFoldersRequest has unknown keys!", "")
    }
    return json.Unmarshal([]byte(s), &r)
}

type ListFoldersResponse struct {
    *ksyunhttp.BaseResponse
    Name *string `json:"Name" name:"Name"`
    Desc *string `json:"Desc" name:"Desc"`
    Id *string `json:"Id" name:"Id"`
    CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
    Level *int `json:"Level" name:"Level"`
    ParentId *int `json:"ParentId" name:"ParentId"`
    Num *int `json:"Num" name:"Num"`
	SonFolder []struct {
		Name *string `json:"Name"`
		Desc *string `json:"Desc"`
		Id *string `json:"Id"`
		CreatedTime *string `json:"CreatedTime"`
		Level *int `json:"Level"`
		ParentId *string `json:"ParentId"`
		Num *int `json:"Num"`
		SonFolder []struct {
			} `json:"SonFolder"`
		} `json:"SonFolder"`
    RequestId *string `json:"RequestId" name:"RequestId"`
}

func (r *ListFoldersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListFoldersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

