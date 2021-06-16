package base

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				context.AbortWithStatusJSON(400, gin.H{"error": e})
			}
		}()
		context.Next()
	}
}

func Error(err error, message ...string) {
	if err == nil {
		return
	} else {
		errMessage := err.Error()
		if len(message) > 0 {
			errMessage = message[0]
		}
		panic(errMessage)
	}
}
