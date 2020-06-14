/**
 * @Author: Hongker
 * @Description: 接口启动程序
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2020/6/14 10:42
 */

package main

import (
	"github.com/ebar-go/ego"
	"github.com/ebar-go/ego/utils/secure"
	"github.com/ebar-go/gateway/http/backend"
	"github.com/ebar-go/gateway/http/frontend"
)

func main() {
	server := ego.HttpServer()

	// 加载前台路由
	frontend.LoadRoute(server.Router)
	// 加载后台路由
	backend.LoadRoute(server.Router)

	secure.Panic(server.Start())
}
