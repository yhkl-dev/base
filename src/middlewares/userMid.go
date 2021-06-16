package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserMid struct {
}

func NewUserMid() *UserMid {
	return &UserMid{}
}

func (s *UserMid) OnRequest(context *gin.Context) error {
	fmt.Println("new middleware")
	fmt.Println(context.Query("name"))
	return nil
}
