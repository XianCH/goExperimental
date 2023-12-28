package pojo

import "github.com/x14n/goExperimental/gin/database"

type User struct {
	Id       int    `json:"UserId"`
	Username string `json:"Username"`
	Passwd   string `json:"Passwd"`
	Email    string `json:"Email"`
}

func FindAllUser() []User {
	var users []User
	database.DB.Find(&users)
	return users
}

func FindByUserId(userId string) User {
	var user User
	database.DB.Where("id = ?", userId).First(&user)
	return user
}

func CreateUser(user User) User {
	database.DB.Create(&user)
	return user
}

func DeleteUser(userId int) User {
	user := User{}
	database.DB.Where("id =?", userId).Delete(&user)
	return user
}
