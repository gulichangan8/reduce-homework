package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("sign/*")

	r.GET("/user", func(c *gin.Context) {
		c.HTML(200, "login1.html", gin.H{})
	})

	r.POST("/login1", func(c *gin.Context) {
		username := c.PostForm("username")
		c.SetCookie("username", username, 3600, "/",
			"localhost", false, true)
		c.HTML(200, "turn1.html", gin.H{})
	})

	r.POST("/next", func(c *gin.Context) {
		cookie, _ := c.Cookie("username")
		if cookie == "" {
			c.String(200, "用户名不能为空")
		} else {
			c.String(200, "登陆成功")
		}
	})
	r.Run()
}
