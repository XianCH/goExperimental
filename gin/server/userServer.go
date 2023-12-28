package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x14n/goExperimental/gin/pojo"
)

var userList = []pojo.User{}

// GetUser
func FindAllUser(c *gin.Context) {
	users := pojo.FindAllUser()
	c.JSON(http.StatusOK, users)
}

//PostUser

func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error : %v"+err.Error())
		return
	}
	userList = append(userList, user)
	c.JSON(http.StatusOK, "successfull")
}

func FindUsreById(c *gin.Context) {
	user := pojo.FindByUserId(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("user->", user)
	return
}

func CreateUser(c *gin.Context) {

}
