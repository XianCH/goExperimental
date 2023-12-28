package httpjwt

import (
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginInput := &LoginInput{}
	if err := loginInput.ValidLoginInput(c); err != nil {
		ResponseError(c, -1, err)
		return
	}

	fmt.Println(loginInput)

	loginResult, err := LoginService(loginInput.UserName, loginInput.Password)
	if err != nil {
		if err.Error() == "record not found" {
			ResponseError(c, 500, errors.New("该用户不存在"))
			return
		} else {
			ResponseError(c, 500, errors.New("登录错误"))
			return
		}
	}

	// fmt.Println(loginResult.User)

	ResponseSuccess(c, loginResult)
}

func UserList(c *gin.Context) {
	var user User
	claims := c.MustGet("claims").(*CustomClaims)
	log.Println(claims)
	users, err := user.ListUsers(claims.Name)
	if err != nil {
		log.Fatal(err)
	}
	ResponseSuccess(c, users)
}
