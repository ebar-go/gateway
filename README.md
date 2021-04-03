# gateway
go http gateway

## 系统架构
![架构图](http://processon.com/chart_image/5c3b0847e4b048f108cdf001.png)

## 概念说明
### Manager Node(管理节点)
提供上游服务的管理操作，并将变更信息同步到多台API节点。
### Api Node (API节点)
接收应用层的API请求，并将请求转发到上游服务。
### Upstream Service 上游服务
上游服务包含服务名称，路由组信息，服务地址及权重等信息。

## 运行
```
cd gateway
GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app cmd/http/main.go
./app
```

## 配置
- 添加upstream
- 添加endpoint