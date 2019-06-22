package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/lottery/wheel/bootstrap"
)


var port = 8080

func newApp() *bootstrap.Bootstrapper {
	// 初始化应用
	app := bootstrap.New("Go 抽奖系统"， "一凡Sir")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)

	return app
}

func main() {
}