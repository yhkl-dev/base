package classes

import (
	"base/src/base"
	"base/src/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserClass struct {
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (s *UserClass) UserTest(context *gin.Context) string {
	return "abc"
	// return func(context *gin.Context) {
	// 	context.JSON(200, gin.H{
	// 		"result": "success",
	// 	})
	// }UserList
}

func (s *UserClass) UserList(context *gin.Context) base.Models {
	users := []*models.UserModel{
		{UserID: 1001, UserName: "yhkl"},
		{UserID: 1002, UserName: "x"},
	}

	return base.MakeModels(users)
	// return &models.UserModel{UserID: 100, UserName: "yhkl"}
}

func (s *UserClass) UserDetail(context *gin.Context) base.Model {
	user := models.NewUserModel()
	err := context.BindUri(user)
	fmt.Println(user)
	base.Error(err, "invalid user id")
	return user
	// return &models.UserModel{UserID: 100, UserName: "yhkl"}
}

func (s *UserClass) Build(base *base.Base) {
	base.
		Handle("GET", "/user", s.UserList).
		Handle("GET", "/user/:id", s.UserDetail)
}
