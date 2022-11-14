package main

import (
	"github.com/gin-gonic/gin"
	"path"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("sign/*")

	r.GET("/user", func(c *gin.Context) {
		c.HTML(200, "register2.html", gin.H{})
	})

	r.POST("/register", func(c *gin.Context) {
		c.HTML(200, "turn2.html", gin.H{})
		file1, _ := c.FormFile("username")
		file2, _ := c.FormFile("password")
		dst1 := path.Join("./file/username", file1.Filename)
		dst2 := path.Join("./file/password", file2.Filename)
		c.SaveUploadedFile(file1, dst1)
		c.SaveUploadedFile(file2, dst2)
	})

	r.POST("/login", func(c *gin.Context) {
		c.HTML(200, "login3.html", gin.H{})
	})

	r.POST("/next", func(c *gin.Context) {
		c.String(200, "登陆成功")
	})
	r.Run()
}
