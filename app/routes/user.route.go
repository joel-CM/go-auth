package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joel-CM/go-auth/app/controller"
)

func UserRoute(app *gin.Engine) {
	api := app.Group("/api")
	user := api.Group("/user")

	user.GET("/", controller.GetAllUsers)
	user.POST("/register", controller.UserRegister)
}
