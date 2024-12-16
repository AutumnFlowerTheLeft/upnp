# Go-NAT-PMP
使用 *狗浪* 开发的UPNP，NAT-PMP

为青联互联RPC通信框架插件，即XoRPC UPNP插件。

XoRPC : https://github.com/qiaoliangXgamemode/XoRPC
# 下载包

```
go get github.com/AutumnFlowerTheLeft/upnp@v1.0.0
```

# Usage
创建一个 *main.go* 文件写入以下内容：
```
package main

import (
	"github.com/AutumnFlowerTheLeft/upnp"
)

func main() {
	gwy, e := upnp.Getgateway()
	if e != nil {
		panic(e)
	}
	response := &upnp.AddPortMapping{
		Getgateway:         gwy,
		Protocol:           "tcp",
		Port:               5005,
		MappedExternalPort: 0,
		TimeOut:            3600,
	}
	upnp.OpenPortMapping(response)
}
```
# Example
## Getgateway 网关
路由器或者网卡网关获取
```
gwy, e := upnp.Getgateway()
```


## AddPortMapping 添加映射端口
```
Getgateway 网关
Protocol 网络协议 tcp/udp
Port 端口
MappedExternalPort 外网端口（通常运营商会随机）
TimeOut  映射停留时间
```
# 最后
查看路由器后台 或者 opwrt 后台，UPNP-PMP 端口是否打开。
