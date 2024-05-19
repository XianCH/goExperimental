package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/x14n/goExperimental/gin-crud-api/v3/api/data"
	"github.com/x14n/goExperimental/gin-crud-api/v3/utils"
)

func TestApi(c *gin.Context) {
	var form data.LoginForm
	if err := c.ShouldBind(&form); err != nil {
		utils.ErrorWithMsg(c, err.Error())
	}
	fmt.Println(form)
	utils.OkWithData(c, form)
}

func CreateEmailCode(c *gin.Context) {

}

func ChackEmail(c *gin.Context) {

}

func RegisterUser(c *gin.Context) {

}

func Login(c *gin.Context) {

}

func GetOnlineUser(c *gin.Context) {

}

func GetUserById(c *gin.Context) {

}

func UploadFile(c *gin.Context) {

}

func GetFileDetilById(c *gin.Context) {

}
