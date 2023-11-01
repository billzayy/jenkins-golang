package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"status": "Successful", "data": "Nothing"})
}
