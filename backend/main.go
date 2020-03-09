package main

import (
	"firstTasteIris/backend/config"
	"firstTasteIris/backend/logs"
	"firstTasteIris/backend/routes"
	"fmt"
	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
)

func main() {
	f := logs.NewLog()
	defer logs.CloseLog(f)

	irisApplication := NewApp()

	err := irisApplication.Run(iris.Addr(config.GetAppUrl()), iris.WithConfiguration(config.GetIrisConf()))
	if err != nil {
		color.Red(fmt.Sprintf("项目运行结束: %v", err))
	}
}

func NewApp() *iris.Application {
	app := iris.New()

	app.Logger().SetLevel(config.GetAppLoggerLevel()) // 从配置中读取日志级别 并设置

	if deployment := config.GetDeployment();deployment == "develop" { // 开发模式
		app.RegisterView(iris.HTML("../front", ".html").Reload(true))// 增加了静态文件重载
	} else {
		app.RegisterView(iris.HTML("../front", ".html"))
	}

	app.HandleDir("/assets", "../front/assets")

	routes.Register(app) //注册路由

	return app
}
