package router

import (
	"github.com/gin-gonic/gin"
	"rakamin.com/final-task/controllers"
	"rakamin.com/final-task/middlewares"
)

func SetupUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.Register)
		userGroup.POST("/login", controllers.Login)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
		userGroup.GET("/active", middlewares.RequireAuth, controllers.GetUserLogin)
	}

	r.GET("/", controllers.PostsIndex)
}
