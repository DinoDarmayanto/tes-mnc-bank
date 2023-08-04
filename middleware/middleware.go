package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Printf("URL %v [%v] dipanggil pada waktu : [%v]\n", ctx.Request.URL.Path, ctx.Request.Method, time.Now())
		ctx.Next()
		fmt.Printf("URL %v [%v] selesai dipanggil pada waktu : [%v]\n", ctx.Request.URL.Path, ctx.Request.Method, time.Now())
	}
}
