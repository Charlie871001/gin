package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Charlie")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	r.Run()
}
