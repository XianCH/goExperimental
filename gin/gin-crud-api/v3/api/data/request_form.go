package data

type LoginForm struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required,min=6"`
}
