package main

import (
	"github.com/gin-gonic/gin"
	"rakamin.com/final-task/database"
	"rakamin.com/final-task/middlewares"
	"rakamin.com/final-task/router"
)

func init() {
	middlewares.LoadEnvVariables()
	database.ConnectoDB()
}

func main() {
	r := gin.Default()

	router.SetupUserRoutes(r)
	router.SetupPhotoRoutes(r)

	r.Run() // listen and serve on http://localhost:3000/
}
