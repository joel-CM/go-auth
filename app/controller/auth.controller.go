package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetResource(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "successful access to resource :)"})
}
