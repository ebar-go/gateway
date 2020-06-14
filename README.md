# gateway
go http gateway

## 系统架构
1.一个master node,运行backend服务
2.多个worker node 运行frontend服务,也就是转发http请求。
3.一个node里有多个pod,一个pod里运行多个相同的container