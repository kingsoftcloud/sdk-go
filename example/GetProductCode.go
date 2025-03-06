package main

import (
	"fmt"
	bill "github.com/kingsoftcloud/sdk-go/v2/ksyun/client/bill/v20180601"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common"
	"github.com/kingsoftcloud/sdk-go/v2/ksyun/common/profile"
)

func main() {
	credential := common.NewCredential(
		"KSYUN_AK_HERE",
		"KSYUN_SK_HERE",
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "GET"
	cpf.HttpProfile.ReqTimeout = 60
	cpf.HttpProfile.Endpoint = "bill.api.ksyun.com"
	client, _ := bill.NewClient(credential, "cn-beijing-6", cpf)

	request := bill.NewGetProductCodeRequest()

	response := client.GetProductCode(request)

	fmt.Printf("%s", response)
}
