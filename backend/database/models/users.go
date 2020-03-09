package models

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/jameskeane/bcrypt"
	"time"

	"firstTasteIris/backend/database"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Name      string `gorm:"not null VARCHAR(191)"`
	Username  string `gorm:"unique;VARCHAR(191)"`
	Password  string `gorm:"not null VARCHAR(191)"`
	TenancyID uint
}

func NewUser(id uint, username string) *User {
	return &User{
		Model: gorm.Model{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: username,
	}
}

func (u *User) GetUserByUsername() {
	IsNotFound(database.GetGdb().Where("username = ?", u.Username).First(u).Error)
}

/**
 * 判断用户是否登录
 * @method CheckLogin
 * @param  {[type]}  id       int    [description]
 * @param  {[type]}  password string [description]
 */
func (u *User) CheckLogin(password string) (*Token, bool, string) {
	if u.ID == 0 {
		return nil, false, "用户不存在"
	} else {
		if ok := bcrypt.Match(password, u.Password); ok {
			token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
				"iat": time.Now().Unix(),
			})
			tokenString, _ := token.SignedString([]byte("HS2JDFKhu7Y1av7b"))

			oauthToken := new(OauthToken)
			oauthToken.Token = tokenString
			oauthToken.UserId = u.ID
			oauthToken.Secret = "secret"
			oauthToken.Revoked = false
			oauthToken.ExpressIn = time.Now().Add(time.Hour * time.Duration(1)).Unix()
			oauthToken.CreatedAt = time.Now()

			response, err := oauthToken.OauthTokenCreate()
			if err != nil {
				return nil, false, "登陆失败"
			}

			return response, true, "登陆成功"
		} else {
			return nil, false, "用户名或密码错误"
		}
	}
}

/**
* 用户退出登陆
* @method UserAdminLogout
* @param  {[type]} ids string [description]
 */
func UserAdminLogout(userId uint) bool {
	ot := OauthToken{}
	ot.UpdateOauthTokenByUserId(userId)
	return ot.Revoked
}
