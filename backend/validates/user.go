package validates

type LoginRequest struct {
	Username string `json:"username" validate:"required,gte=5,lte=50" comment:"user name"`
	Password string `json:"password" validate:"required" comment:"password"`
}

// 登陆表单验证
func (alr *LoginRequest) Valid() string {
	return BaseValid(alr)
}