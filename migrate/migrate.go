package main

import (
	"rakamin.com/final-task/database"
	"rakamin.com/final-task/middlewares"
	"rakamin.com/final-task/models"
)

func init() {
	database.ConnectoDB()
	middlewares.LoadEnvVariables()
}

func main() {
	database.DB.AutoMigrate(&models.User{}, &models.Photo{})
}
