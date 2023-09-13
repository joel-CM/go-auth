package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joel-CM/go-auth/app/db"
	"github.com/joel-CM/go-auth/app/models"
	"github.com/joel-CM/go-auth/app/routes"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := gin.Default()

	// db connection
	db.DBConnect()

	// model migrations
	db.GormDB.AutoMigrate(&models.User{})

	// all routes
	routes.Index(app)

	// listen and serve on 0.0.0.0:3001
	app.Run(":3001")
}
