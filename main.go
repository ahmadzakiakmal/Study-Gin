package main

import (
	"fmt"

	"github.com/ahmadzakiakmal/Study-Gin/initializers"
	"github.com/ahmadzakiakmal/Study-Gin/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.POST("/user", func(c *gin.Context) {
		var body models.User
		err := c.ShouldBindJSON(&body)
		if err != nil {
			fmt.Println(err)
			c.JSON(415, map[string]interface{}{
				"error": "invalid JSON",
			})
			return
		}
		c.JSON(200, map[string]interface{}{
			"name": body.Name,
			"age":  body.Age,
		})
	})

	r.Run()
}
