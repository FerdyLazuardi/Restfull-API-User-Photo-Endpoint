package router

import (
	"github.com/gin-gonic/gin"
	"rakamin.com/final-task/controllers"
	"rakamin.com/final-task/middlewares"
)

func SetupPhotoRoutes(r *gin.Engine) {
	photoGroup := r.Group("/photo")
	{
		photoGroup.GET("/", middlewares.RequireAuth, controllers.GetPhoto)
		photoGroup.POST("/", middlewares.RequireAuth, controllers.AddPhoto)
		photoGroup.PUT("/:id", middlewares.RequireAuth, controllers.UpdatePhoto)
		photoGroup.DELETE("/:id", middlewares.RequireAuth, controllers.DeletePhoto)
	}
}
