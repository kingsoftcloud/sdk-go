<h1 align="center">Kingsoft cloud SDK for Go</h1>

# 简介

欢迎使用金山云开发者工具套件（SDK），此SDK是金山云API平台的配套开发工具。</br>
当前SDK版本 1.0.6

# 依赖环境

1. Go 1.17 版本及以上。
2. 在金山云控制台页面获取密钥 AccessKey 和 SecretKey，请务必妥善保管，或者使用更安全的临时安全凭证。

# 获取安装

注意：此安装方式仅支持使用 **Go Modules** 模式进行依赖管理，即环境变量 `GO111MODULE=auto`或者`GO111MODULE=on`,
并且在您的项目中执行了 `go mod init xxx`.

如果您使用 GOPATH, 请参考下节： 全部安装

安装SDK包：

    ```bash
    go get -v -u github.com/kingsoftcloud/sdk-go/ksyun/common
    ```

## 通过源码安装

前往代码托管地址 [Github](https://github.com/kingsoftcloud/) 下载最新代码，解压后安装到项目目录下。

# 快速开始

每个接口都有一个对应的 Request 结构和一个 Response 结构。例如访问控制中查询子用户列表接口ListUsers,
有对应的请求结构体ListUsersRequest 和返回结构体 ListUsersResponse 。

下面以访问控制查询子用户列表接口为例，介绍 SDK 的基础用法。

##  

Demo

```go

package main

import (
"fmt"
"github.com/kingsoftcloud/sdk-go/ksyun/common"
"github.com/kingsoftcloud/sdk-go/ksyun/common/errors"
"github.com/kingsoftcloud/sdk-go/ksyun/common/profile"
iam "github.com/kingsoftcloud/sdk-go/ksyun/iam/v20151101"
)

func main() {
   credential := common.NewCredential("AccessKey", "SecretKey")

   cpf := profile.NewClientProfile()
   cpf.HttpProfile.ReqMethod = "POST"
   //设置超时时间  可不设置
   cpf.HttpProfile.ReqTimeout = 10
   //请求域名
   cpf.HttpProfile.Endpoint = "iam.api.ksyun.com"
   client, _ := iam.NewClient(credential, "cn-beijing-6", cpf)

   request := iam.NewListUsersRequest()
   request.MaxItems = common.Int64Ptr(1)
  
   response, err := client.ListUsers(request)

   if _, ok := err.(*errors.KsyunSDKError); ok {
      fmt.Printf("An API error has returned: %s", err)
      return
   }

   if err != nil {
      panic(err)
   }

   fmt.Printf("%s", response.ToJsonString())
}

```

# 相关配置

**如无特殊需要，建议您使用默认配置。**

在创建客户端前，如有需要，您可以通过修改`profile.ClientProfile`中字段的值进行一些配置。

```go
// 【非必需】 实例化一个客户端配置对象，可以指定超时时间等配置
cpf := profile.NewClientProfile()
```

具体的配置项说明如下：

## 请求方式

SDK默认使用POST方法。 如果你一定要使用GET方法，可以在这里设置。**GET方法无法处理一些较大的请求**。

```go
cpf.HttpProfile.ReqMethod = "POST"
```

## 超时时间

SDK有默认的超时时间，如非必要请不要修改默认设置。
如有需要请在核实后获取最新的默认值并进行代码配置，单位(秒)。

```go
cpf.HttpProfile.ReqTimeout = 10
```

## 指定域名

SDK会自动指定域名。通常是不需要代码指定域名，但是如果你访问的域名需要特殊指定，则进行代码配置

```go
cpf.HttpProfile.Endpoint = "iam.api.ksyun.com"
```

## 签名方式

SDK默认用 `HMAC-SHA256` 进行签名，它更安全但是会轻微降低性能。

```go
cpf.SignMethod = "HMAC-SHA256"
```

## DEBUG模式

DEBUG模式会打印更详细的日志，当您需要进行详细的排查错误时可以开启。  
默认为 `false`

```go
cpf.Debug = true
```
