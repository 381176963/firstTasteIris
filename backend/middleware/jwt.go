package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

/**
 * 验证 jwt
 * @method JwtHandler
 */
func JwtHandler() *jwt.Middleware {
	var mySecret = []byte("HS2JDFKhu7Y1av7b")
	return jwt.New(jwt.Config{
		ErrorHandler: func(ctx iris.Context, err error) {
			if err == nil {
				return
			}

			ctx.StopExecution()
			ctx.StatusCode(iris.StatusUnauthorized)
			_, _ = ctx.JSON(struct{
				Status bool `json:"status"`
				Msg    interface{} `json:"msg"`
				Data   interface{} `json:"data"`
			}{Status: false, Data: struct{}{}, Msg: "你的token不正确"})
		},
		//Extractor: jwt.FromParameter("token"),
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}
