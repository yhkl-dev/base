package base

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

var ResponderList []Responder

func init() {
	ResponderList = []Responder{new(StringResponder), new(ModelResponder), new(ModelsResponder)}
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

func Convert(handler interface{}) gin.HandlerFunc {
	hRef := reflect.ValueOf(handler)
	for _, r := range ResponderList {
		rRef := reflect.ValueOf(r).Elem()
		if hRef.Type().ConvertibleTo(rRef.Type()) {
			rRef.Set(hRef)
			return rRef.Interface().(Responder).RespondTo()
		}
	}
	return nil
}

type StringResponder func(*gin.Context) string

func (s StringResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(200, s(context))
	}
}

type ModelResponder func(*gin.Context) Model

func (s ModelResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, s(context))
	}
}

type ModelsResponder func(*gin.Context) Models

func (s ModelsResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Content-Type", "application/json")
		context.Writer.WriteString(string(s(context)))
	}
}
