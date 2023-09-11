package routes

import "github.com/gin-gonic/gin"

func Index(app *gin.Engine) {
	UserRoute(app)
	AuthRoute(app)
}
