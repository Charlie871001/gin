package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1、创建路由
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
	r.Run()
}
