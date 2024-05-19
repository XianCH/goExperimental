package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateForm(c *gin.Context, form interface{}) error {
	// 解析表单数据
	if err := c.Bind(form); err != nil {
		return err
	}

	// 使用验证器对表单数据进行验证
	validate := validator.New()
	if err := validate.Struct(form); err != nil {
		return err
	}

	// 如果验证通过，则返回 nil
	return nil
}
