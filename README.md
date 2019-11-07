## zinglabs-proxy

 研发跳板机，初步用于方便观看连接了阿里云的VPN的机器，端口转发
 
*golang 实现多平台交叉编译：*

### Mac 下编译 Linux 和 Windows 64位可执行程序

```shell
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```
  
### Linux 下编译 Mac 和 Windows 64位可执行程序

```shell
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

### Windows 下编译 Mac 和 Linux 64位可执行程序

```shell
    SET CGO_ENABLED=0
    SET GOOS=darwin
    SET GOARCH=amd64
    go build main.go
```

去 mian 目录打包，打包完成取出相应平台的可执行文件，双击执行「目前写死 105」

## 目前的代理：

```
    devEs     = "dev.es"
	devConsul = "dev.consul"
	devKong   = "dev.kong"
	devKonga  = "dev.konga"
	devMinio  = "dev.minio"

	prodEs     = "prod.es"
	prodConsul = "prod.consul"
	prodKong   = "prod.kong"
	prodKonga  = "prod.konga"
	prodMinio  = "prod.minio"

	globElk = "glob.elk"
```


## 用法：

 目前我们的跳板机是 `192.168.18.153`，但是 windows 远程很容易被抢占，所以使用 golang 协程来转发。
 
 浏览器输入：
 
 ```shell
 http://192.168.18.153/dev.es #测试 es
 http://192.168.18.153/glob.elk #所有环境的 elk 日志查询
 http://192.168.18.153/prod.es #生产 es 慎重！！！！
 ```
 
# 所有的 `prod` 都要慎重！！

未来想支持：内网代理，命令行模式，可配置的远程列表，以及类似 SS 的代理模式「需要电脑设置 HTTP 代理」

## html 、css 等静态文件打成二进制

https://github.com/GeertJohan/go.rice

html 有改动需要重新执行一下 rice。

