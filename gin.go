package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 建立Gin實例
	r := gin.Default()

	// 設置路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// 帶有路由參數的路由
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	// 帶有中間件的路由
	r.GET("/middleware", middleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is a route with middleware.",
		})
	})

	// 模板渲染路由
	r.GET("/template", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Gin Template",
		})
	})

	// 啟動Gin
	r.Run(":8080")
}

// 定義中間件
func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("middleware_message", "This is a middleware message.")
		c.Next()
	}
}
