package main

import (
	"github.com/Bewok19/first-api/config"
	"github.com/Bewok19/first-api/modules/user"
	"github.com/Bewok19/first-api/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB() // connect toa database

	userRepository := repositories.NewUserRepository(config.DB)
	userService := user.NewUseCase(userRepository)
	userController := user.NewUserController(userService)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	user := router.Group("/user")
	user.GET("all", userController.FindAll)

	router.Run(":8080")
}
