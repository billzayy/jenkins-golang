package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserServices(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Successful", "data": "Nothing"})
}
