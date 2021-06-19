package base

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Base struct {
	*gin.Engine
	g           *gin.RouterGroup
	beanFactory *BeanFactory
}

func Ignite() *Base {
	g := &Base{Engine: gin.New(), beanFactory: NewBeanFactory()}
	g.Use(ErrorHandler())
	g.beanFactory.setBean(InitConfig())
	return g
}

func (b *Base) Launch() {
	// config := InitConfig()
	// b.Run(fmt.Sprintf(":%d", config.Server.Port))
	var port int32 = 8080
	if config := b.beanFactory.GetBean(new(SysConfig)); config != nil {
		port = config.(*SysConfig).Server.Port
	}
	b.Run(fmt.Sprintf(":%d", port))
}

func (b *Base) Attach(f Fairing) *Base {
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

func (b *Base) Beans(beans ...interface{}) *Base {
	// b.props = append(b.props, beans...)
	b.beanFactory.setBean(beans...)
	return b
}

func (b *Base) Handle(httpMethod, relativePath string, handlers interface{}) *Base {
	// b.g.Handle(httpMethod, relativePath, handlers...)
	if h := Convert(handlers); h != nil {
		b.g.Handle(httpMethod, relativePath, h)
	}
	return b
}

func (b *Base) Mount(group string, classes ...IClass) *Base {
	// g := b.Group(group)
	b.g = b.Group(group)
	for _, class := range classes {
		class.Build(b)
		b.beanFactory.inject(class)
	}
	return b
}

// // get props
// func (b *Base) getProp(t reflect.Type) interface{} {
// 	for _, p := range b.props {
// 		if t == reflect.TypeOf(p) {
// 			return p
// 		}
// 	}
// 	return nil
// }

// func (b *Base) setProp(class IClass) {
// 	vClass := reflect.ValueOf(class).Elem()
// 	vClassT := reflect.TypeOf(class).Elem()
// 	for i := 0; i < vClass.NumField(); i++ {
// 		f := vClass.Field(i)
// 		if !f.IsNil() || f.Kind() != reflect.Ptr {
// 			continue
// 		}
// 		if p := b.getProp(f.Type()); p != nil {
// 			f.Set(reflect.New(f.Type().Elem()))
// 			f.Elem().Set(reflect.ValueOf(p).Elem())
// 			if IsAnnotation(f.Type()) {
// 				p.(Annotation).SetTag(vClassT.Field(i).Tag)
// 			}
// 		}
// 	}
// }
