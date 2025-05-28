package main

import (
	"fmt"
	iam "github.com/kingsoftcloud/sdk-go/v2/ksyun/client/iam/v20151101"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	_ "github.com/kingsoftcloud/sdk-go/v2/ksyun/common/errors"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

func main() {
	credential := common.NewCredential(
		"KSYUN_AK_HERE",
		"KSYUN_SK_HERE",
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	//设置超时时间  可不设置
	cpf.HttpProfile.ReqTimeout = 10
	//请求域名  不要加http://
	cpf.HttpProfile.Endpoint = "iam.api.ksyun.com"
	client, _ := iam.NewClient(credential, "cn-shanghai-3", cpf)

	request := iam.NewListUsersRequest()
	// 	request.MaxItems = common.StringPtr("zone")
	// 	request.String = common.StringPtr("zone")
	// 	request.Users = []string{"liyukun_01"}
	// 	request.Policies = []string{"krn:ksc:resourcemanager::ksc:control-policy/LY_sys"}
	// 	request.Filter = []*iam.Filter{
	// 		&iam.Filter{
	// 			Name:   common.StringPtr("zone"),
	// 			Values: common.StringPtrs([]string{"ap-guangzhou-1"}),
	// 		},
	// 	}

	//支持设置请求头
	//request.SetHeaders(map[string]string{
	//	"KEY": "VALUE",
	//})
	response := client.ListUsers(request)

	fmt.Printf("%s", response)
}
