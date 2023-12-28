package httpjwt

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type LoginInput struct {
	UserName string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

func (o *LoginInput) ValidLoginInput(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	//翻译库
	validate := validator.New()
	if err := validate.Struct(o); err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}
