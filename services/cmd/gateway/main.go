package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"services/internal/gateway"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Database Error")
	}

	if err = db.AutoMigrate(&gateway.User{}); err != nil {
		panic("Database Migration Error")
	}

	gateway.InitializeRoutes(r, db)

	err = r.Run(":8080")
	if err != nil {
		panic("Error starting server")
	}
}
