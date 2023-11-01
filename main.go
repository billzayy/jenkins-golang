package main

import (
	"os"

	"github.com/billzayy/jenkins-golang/services"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	api := app.Group("/api")
	{
		api.GET("/get-all-user", services.GetUser)
	}

	app.Run(":" + os.Getenv("PORT"))
}
