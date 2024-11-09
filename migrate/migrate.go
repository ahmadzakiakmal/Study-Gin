package main

import (
	"github.com/ahmadzakiakmal/Study-Gin/initializers"
	"github.com/ahmadzakiakmal/Study-Gin/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
