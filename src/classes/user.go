package classes

import (
	"base/src/base"

	"github.com/gin-gonic/gin"
)

type UserClass struct {
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (s *UserClass) UserList() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"result": "success",
		})
	}
}

func (s *UserClass) Build(base *base.Base) {
	base.Handle("GET", "/user", s.UserList())
}
