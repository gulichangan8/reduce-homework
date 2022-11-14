package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	user := make(map[string]string)

	r := gin.Default()
	r.LoadHTMLGlob("sign/*")

	r.GET("/user", func(c *gin.Context) {
		c.HTML(200, "register1.html", gin.H{})
	})

	r.POST("/register", func(c *gin.Context) {
		remember(c, user)
	})

	r.POST("/login", func(c *gin.Context) {
		c.HTML(200, "login2.html", gin.H{})
	})

	r.POST("/next", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if user[username] == password {
			c.String(200, "登陆成功")
		} else {
			c.String(200, "登陆失败")
		}
	})
	r.Run()
}

func remember(c *gin.Context, user map[string]string) map[string]string {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user[username] = password
	c.HTML(200, "turn2.html", gin.H{})
	return user
}
