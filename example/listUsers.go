package main

import (
	"fmt"
	iam "github.com/kingsoftcloud/sdk-go/ksyun/client/iam/v20151101"
	"github.com/kingsoftcloud/sdk-go/ksyun/common"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
	"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
)

func main() {
	ak := "AKLT480W7RcdQKS_D33a2K11yw"
	sk := "OHzdodDebcRVcDqSdefxbjCix60LeLUbLKA3rxyOvu3uauDyr0WC/FIc6weB/YXYBQ=="

	credential := common.NewCredential(
		ak,
		sk,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	//设置超时时间  可不设置
	cpf.HttpProfile.ReqTimeout = 10
	//请求域名  不要加http://
	cpf.HttpProfile.Endpoint = "10.111.94.163"
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
	response, err := client.ListUsers(request)

	if _, ok := err.(*errors.KsyunSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}

	if err != nil {
		fmt.Printf("An API error has returned: %s", err)
	}

	fmt.Printf("%s", response.ToJsonString())
}
