package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joel-CM/go-auth/app/controller"
)

func UserRoute(app *gin.Engine) {
	api := app.Group("/api")   // api path group
	user := api.Group("/user") // user path group

	user.POST("/register", controller.UserRegister) // user register
	user.POST("/signin", controller.UserSignIn) // user signin/login
}
