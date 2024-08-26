package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunRoutes() error {
	router := gin.Default()

	router.GET("/ping", PingPong)
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	userG := router.Group("/users")
	{
		userG.GET("", GetAllUsers)
		userG.GET("/:id", GetUserByID)
		userG.POST("", CreateUser)
		userG.PUT("/:id")
		userG.DELETE("/:id")
		userG.PATCH("/:id")
	}

	taskG := router.Group("/tasks")
	{
		taskG.GET("")
		taskG.POST("")
		taskG.GET("/:id")
		taskG.PUT("/:id")
		taskG.DELETE("/:id")
	}

	err := router.Run(":8181")
	if err != nil {
		return err
	}

	return nil
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
