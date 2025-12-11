package v20210320
import (
	"encoding/json"
	ksyunhttp "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/http"
)


type CreateFolderRequest struct {
	*ksyunhttp.BaseRequest
	ParentId *string `json:"ParentId,omitempty" name:"ParentId"`
	Name     *string `json:"Name,omitempty" name:"Name"`
	Desc     *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateFolderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	Name     *string `json:"Name,omitempty" name:"Name"`
	Desc     *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *UpdateFolderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	Search   *string `json:"Search,omitempty" name:"Search"`
	Page     *int    `json:"Page,omitempty" name:"Page"`
	PageSize *int    `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *ListAccountsForParentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListAccountsForParentResponse struct {
	*ksyunhttp.BaseResponse
	Members []struct {
		CreatedTime     *string `json:"CreatedTime" name:"CreatedTime"`
		Name            *string `json:"Name" name:"Name"`
		AdminPermission *int    `json:"AdminPermission" name:"AdminPermission"`
		UserId          *int    `json:"UserId" name:"UserId"`
		UserName        *string `json:"UserName" name:"UserName"`
		UserType        *int    `json:"UserType" name:"UserType"`
		Phone           *string `json:"Phone" name:"Phone"`
		FolderId        *string `json:"FolderId" name:"FolderId"`
	} `json:"Members"`
	Count     *int    `json:"Count" name:"Count"`
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
	Ids          *string `json:"Ids,omitempty" name:"Ids"`
	FromFolderId *string `json:"FromFolderId,omitempty" name:"FromFolderId"`
	ToFolderId   *string `json:"ToFolderId,omitempty" name:"ToFolderId"`
}

func (r *MoveAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	MemberId       *int    `json:"MemberId,omitempty" name:"MemberId"`
	NewDisplayName *string `json:"NewDisplayName,omitempty" name:"NewDisplayName"`
	FolderId       *string `json:"FolderId,omitempty" name:"FolderId"`
}

func (r *UpdateAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
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
	PageSize   *int `json:"PageSize,omitempty" name:"PageSize"`
	IsAll      *int `json:"IsAll,omitempty" name:"IsAll"`
}

func (r *ListAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type ListAccountsResponse struct {
	*ksyunhttp.BaseResponse
	Accounts []struct {
		AccountId *int    `json:"AccountId" name:"AccountId"`
		AccountName *string `json:"AccountName" name:"AccountName"`
		DisplayName *string `json:"DisplayName" name:"DisplayName"`
		Type      *string `json:"Type" name:"Type"`
		JoinedTime *string `json:"JoinedTime" name:"JoinedTime"`
	} `json:"Accounts"`
	Pagination struct {
		TotalCount *int    `json:"TotalCount" name:"TotalCount"`
		PageNumber *string `json:"PageNumber" name:"PageNumber"`
		PageSize   *string `json:"PageSize" name:"PageSize"`
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

type ListFoldersResponse struct {
	*ksyunhttp.BaseResponse
	Name        *string `json:"Name" name:"Name"`
	Desc        *string `json:"Desc" name:"Desc"`
	Id          *string `json:"Id" name:"Id"`
	CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
	Level       *int    `json:"Level" name:"Level"`
	ParentId    *int    `json:"ParentId" name:"ParentId"`
	Num         *int    `json:"Num" name:"Num"`
	SonFolder   []struct {
		Name  *string `json:"Name" name:"Name"`
		Desc  *string `json:"Desc" name:"Desc"`
		Id    *string `json:"Id" name:"Id"`
		CreatedTime *string `json:"CreatedTime" name:"CreatedTime"`
		Level *int    `json:"Level" name:"Level"`
		ParentId *string `json:"ParentId" name:"ParentId"`
		Num   *int    `json:"Num" name:"Num"`
		SonFolder []struct {
		} `json:"SonFolder" name:"SonFolder"`
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

