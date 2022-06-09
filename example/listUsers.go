package main

import (
	"fmt"
	iam "github.com/kingsoftcloud/sdk-go/ksyun/client/iam/v20151101"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	_ "github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

func main() {
	credential := common.NewCredential(
		"AKLTJH6a6I8lRKisbTnGmcZbjw",
		"OO+oVgJ34o+dYZyeMqmpvICg1MN9zG5qGnw4CrQ5hwZePc8+lFE3QibmIjsCzZKMxw==",
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
	response := client.ListUsers(request)

	fmt.Printf("%s", response)
}
