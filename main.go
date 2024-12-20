package main

import (
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

	routes.SetupRoutes(g, db)

	g.Run(":8080")
}
