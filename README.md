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

测试环境 、 投产环境 、 生产环境  的 consul kong konga es minio

```
- [测试环境 ES 地址](http://192.168.18.153/dev.es)
- [测试环境 地址](http://192.168.18.153/dev.consul)
- [测试环境 Konga 地址](http://192.168.18.153/dev.konga)
- [测试环境 Minio 地址](http://192.168.18.153/dev.minio)
- [测试环境 consul 地址](http://192.168.18.153/dev.consul)

### 投产

- [投产环境 ES 地址！](http://192.168.18.153/pre.es)
- [投产环境 地址！](http://192.168.18.153/pre.consul)
- [投产环境 Konga 地址！](http://192.168.18.153/pre.konga)
- [投产环境 Minio 地址！](http://192.168.18.153/pre.minio)
- [投产环境 consul 地址！](http://192.168.18.153/pre.consul)

### 生产

- [生产环境 ES 地址！](http://192.168.18.153/prod.es)
- [生产环境 consul 地址！](http://192.168.18.153/prod.consul)
- [生产环境 Konga 地址！](http://192.168.18.153/prod.konga)
- [生产环境 Minio 地址！](http://192.168.18.153/prod.minio)

```


## 用法：

 目前我们的跳板机是 `192.168.18.153`，因为 windows 远程桌面只支持一个用户，很容易被抢占，所以使用 golang 协程来转发。
 
## 阿里云代理列表


### 测试

- [测试环境 ES 地址](http://192.168.18.153/dev.es)
- [测试环境 地址](http://192.168.18.153/dev.consul)
- [测试环境 Konga 地址](http://192.168.18.153/dev.konga)
- [测试环境 Minio 地址](http://192.168.18.153/dev.minio)
- [测试环境 consul 地址](http://192.168.18.153/dev.consul)

### 投产

- [投产环境 ES 地址！](http://192.168.18.153/pre.es)
- [投产环境 地址！](http://192.168.18.153/pre.consul)
- [投产环境 Konga 地址！](http://192.168.18.153/pre.konga)
- [投产环境 Minio 地址！](http://192.168.18.153/pre.minio)
- [投产环境 consul 地址！](http://192.168.18.153/pre.consul)

### 生产

- [生产环境 ES 地址！](http://192.168.18.153/prod.es)
- [生产环境 consul 地址！](http://192.168.18.153/prod.consul)
- [生产环境 Konga 地址！](http://192.168.18.153/prod.konga)
- [生产环境 Minio 地址！](http://192.168.18.153/prod.minio)



- [全局ELK日志地址](http://192.168.18.153:5601/app/kibana#/home?_g=())

- [电商测试环境最新 swagger](http://192.168.18.153/swagger.mall)
# 所有的 生产 都要慎重！！ 

未来想支持：内网代理，命令行模式，可配置的远程列表，以及类似 SS 的代理模式「需要电脑设置 HTTP 代理」

## html 、css 等静态文件打成二进制

https://github.com/GeertJohan/go.rice

html 有改动需要重新执行一下 rice。

