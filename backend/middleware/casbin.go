package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"firstTasteIris/backend/database/models"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12/context"

	"github.com/casbin/casbin/v2"
)


func New(e *casbin.Enforcer) *Casbin {
	return &Casbin{enforcer: e}
}


func (c *Casbin) ServeHTTP(ctx context.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	foobar := value.Claims.(jwt.MapClaims)
	for key, value := range foobar {
		ctx.Writef("%s = %v", key, value)
		fmt.Println(value)
	}

	token := models.OauthToken{}
	token.GetOauthTokenByToken(value.Raw) //获取 access_token 信息
	if token.Revoked || token.ExpressIn < time.Now().Unix() {
		//_, _ = ctx.Writef("token 失效，请重新登录") // 输出到前端
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.StopExecution()
		return
	} else if !c.Check(ctx.Request(), strconv.FormatUint(uint64(token.UserId), 10)) {
		ctx.StatusCode(http.StatusForbidden) // Status Forbidden
		_, _ = ctx.Writef("您没有权限") // 输出到前端
		ctx.StopExecution()
		return
	} else {
		ctx.Values().Set("auth_user_id", token.UserId)
	}

	ctx.Next()
}

// Casbin is the auth services which contains the casbin enforcer.
type Casbin struct {
	enforcer *casbin.Enforcer
}

// Check checks the username, request's method and path and
// returns true if permission grandted otherwise false.
func (c *Casbin) Check(r *http.Request, userId string) bool {
	method := r.Method
	path := r.URL.Path
	fmt.Println("请求者是: ", userId)
	fmt.Println("路径是: ", path)
	fmt.Println("方法是: ", method)
	ok, _ := c.enforcer.Enforce(userId, path, method)
	return ok
}
