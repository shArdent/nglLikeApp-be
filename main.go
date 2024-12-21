package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nglLike/database"
	"github.com/nglLike/routes"
)

func main() {
	db, err := database.ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	g := gin.Default()
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Izinkan semua origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Cache preflight request selama 12 jam
	}))
	routes.SetupRoutes(g, db)

	g.Run(":8080")
}
