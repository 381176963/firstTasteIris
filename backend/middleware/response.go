package middleware

import (
	"firstTasteIris/backend/config"
	"github.com/kataras/iris/v12"
	"time"
)

func AdminAuth(ctx iris.Context) {
	ctx.Header("Api_version", config.GetApiVersion())
	ctx.Header("Response_time", time.Now().Format("2006-01-02 15:04:05"))

	ctx.Next()
}
