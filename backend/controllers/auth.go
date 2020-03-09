package controllers

import (
	"firstTasteIris/backend/database/models"
	"firstTasteIris/backend/libs"
	"firstTasteIris/backend/validates"
	"github.com/kataras/iris/v12"
	"net/http"
)

func UserLogin(ctx iris.Context) {
	loginParameters := new(validates.LoginRequest)

	if err := ctx.ReadJSON(loginParameters); err != nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(ApiResource(false, nil, err.Error()))
		return
	}

	if formErrs := loginParameters.Valid(); len(formErrs) > 0 {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(ApiResource(false, nil, formErrs))
		return
	}

	user := models.NewUser(0, loginParameters.Username)
	user.GetUserByUsername()

	response, status, msg := user.CheckLogin(loginParameters.Password)
	if status {
		ctx.Application().Logger().Infof("%s 登录系统", loginParameters.Username)
	}

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(status, response, msg))
	return
}

func UserLogout(ctx iris.Context) {
	aui := ctx.Values().GetString("auth_user_id")
	uid := uint(libs.ParseInt(aui, 0))
	models.UserAdminLogout(uid)

	ctx.Application().Logger().Infof("%d 退出系统", uid)
	ctx.StatusCode(http.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, nil, "退出"))
}
