package base

import (
	"github.com/gin-gonic/gin"
)

type Base struct {
	*gin.Engine
	g *gin.RouterGroup
}

func Ignite() *Base {
	return &Base{Engine: gin.New()}
}

func (b *Base) Launch() {
	b.Run(":8080")
}

func (b *Base) Attach(f Fairing) *Base {
	// b.Use(f)
	b.Use(func(context *gin.Context) {
		err := f.OnRequest(context)
		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		} else {
			context.Next()
		}
	})
	return b
}

func (b *Base) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *Base {
	b.g.Handle(httpMethod, relativePath, handlers...)
	return b
}

func (b *Base) Mount(group string, classes ...IClass) *Base {
	// g := b.Group(group)
	b.g = b.Group(group)
	for _, class := range classes {
		class.Build(b)
	}
	return b
}
