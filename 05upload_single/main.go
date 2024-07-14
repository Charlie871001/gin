package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片错误")
		}
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})
	r.Run()
}
