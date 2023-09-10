package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joel-CM/go-auth/app/db"
	"github.com/joel-CM/go-auth/app/models"
	"github.com/joel-CM/go-auth/app/routes"
)

func main() {
	app := gin.Default()

	// db connect
	db.DBConnect()

	// migrations
	db.GormDB.AutoMigrate(&models.User{})

	// routes
	routes.Index(app)

	app.Run(":3001") // listen and serve on 0.0.0.0:8080
}
