package main

import (
	"os"

	"github.com/billzayy/jenkins-golang/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := gin.Default()

	api := app.Group("/api")
	{
		api.GET("/get-all-user", services.GetUser)
	}

	app.Run(":" + os.Getenv("PORT"))
}
