package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joel-CM/go-auth/app/controller"
	"github.com/joel-CM/go-auth/app/middlewares"
)

func AuthRoute(app *gin.Engine) {
	api := app.Group("/api")
	auth := api.Group("/auth", middlewares.AuthMiddleware())

	auth.GET("/resource", controller.GetResource)
}
