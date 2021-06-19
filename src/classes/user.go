package classes

import (
	"base/src/base"
	"base/src/models"

	"github.com/gin-gonic/gin"
)

type UserClass struct {
	*base.GormAdapter
	Age *base.Value `prefix:"user.age"`
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (s *UserClass) UserTest(context *gin.Context) string {
	return "abc" + s.Age.String()
}

func (s *UserClass) UserList(context *gin.Context) base.Models {
	users := []*models.UserModel{
		{UserID: 1001, UserName: "yhkl"},
		{UserID: 1002, UserName: "x"},
	}
	return base.MakeModels(users)
}

func (s *UserClass) UserDetail(context *gin.Context) base.Model {
	user := models.NewUserModel()
	err := context.BindUri(user)
	base.Error(err, "invalid user id")
	s.DB.Table("t_users").Where("id = ?", user.UserID).Find(user)
	return user
}

func (s *UserClass) Build(base *base.Base) {
	base.
		Handle("GET", "/userTest", s.UserTest).
		Handle("GET", "/user", s.UserList).
		Handle("GET", "/user/:id", s.UserDetail)
}
