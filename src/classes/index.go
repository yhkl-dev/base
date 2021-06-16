package classes

import (
	"base/src/base"

	"github.com/gin-gonic/gin"
)

type IndexClass struct {
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}

func (s *IndexClass) GetIndex() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"result": "success",
		})
	}
}

func (s *IndexClass) Build(base *base.Base) {
	base.Handle("GET", "/", s.GetIndex())
}
