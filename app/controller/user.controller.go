package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joel-CM/go-auth/app/db"
	"github.com/joel-CM/go-auth/app/models"
)

func GetAllUsers(c *gin.Context) {
	users := []models.User{}
	db.GormDB.Find(&users)
	c.IndentedJSON(200, users)
}

func UserRegister(c *gin.Context) {
	user := models.User{}
	validate := validator.New()

	var err error
	if err = c.ShouldBindJSON(&user); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	if err = validate.Struct(user); err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if result := db.GormDB.Create(&user); result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusConflict, gin.H{"message": result.Error.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "user created!"})
}
