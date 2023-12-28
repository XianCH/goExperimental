package httpjwt

import (
	"fmt"
	"log"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/x14n/goExperimental/gorm/gorm_database"
	"gorm.io/gorm"
)

type User struct {
	Id       int    `form:"id" json:"id" gorm:"PRIMARY_KEY"`
	Name     string `form:"username" json:"username"`
	Email    string `form:"email" json:"email",binding:"required"`
	Password string `form:"password" json:"-",binding:"required"`
}

type LoginResult struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

func LoginService(name string, password string) (result LoginResult, err error) {
	var User User
	fmt.Println(name, password)

	obj := gorm_database.DB.Table("Users").Where("Username=? and Passwd = ?", name, password).First(&User)
	if err = obj.Error; err != nil {
		return
	}
	genToken := GenToken(User)
	return genToken, nil
}

func GenToken(user User) LoginResult {
	j := &JWT{
		[]byte("newtoken"),
	}
	claims := &CustomClaims{
		user.Id,
		user.Name,
		user.Password,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "cfun",
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		return LoginResult{
			User:  user,
			Token: token,
		}
	}

	log.Println(token)
	data := LoginResult{
		User:  user,
		Token: token,
	}
	return data
}

func (u *User) ListUsers(name string) (users []User, err error) {
	query := gorm_database.DB
	if name != "" {
		query = query.Where("Username=?", name)
	}
	err = query.Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}
