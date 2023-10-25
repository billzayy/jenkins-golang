package main

import (
	"github.com/billzayy/jenkins-golang/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Link"},
		AllowAllOrigins:  true,
		AllowCredentials: false,
		MaxAge:           300,
	}))

	api := app.Group("/api")
	{
		api.GET("/get-user", services.GetUser)
	}

	app.Run(":1234")
}
