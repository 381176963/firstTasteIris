package routes

import (
	"firstTasteIris/backend/controllers"
	"firstTasteIris/backend/database"
	"firstTasteIris/backend/middleware"
	"github.com/kataras/iris/v12"
)

func Register(api *iris.Application) {
	api.Post("/admin/login", middleware.AdminAuth, controllers.UserLogin)
	app := api.Party("/admin", middleware.CrsAuth(), middleware.AdminAuth).AllowMethods(iris.MethodOptions)
	{
		casbinMiddleware := middleware.New(database.GetEnforcer())
		app.Use(middleware.JwtHandler().Serve, casbinMiddleware.ServeHTTP) //登录验证

		app.Get("/logout", controllers.UserLogout).Name = "退出"
	}

	api.Get("/", func(ctx iris.Context) { // 代理到前端的首页
		_ = ctx.View("index.html")
	})
}
