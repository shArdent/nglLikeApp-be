package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/nglLike/controllers"
)

func SetupRoutes(g *gin.Engine, db *sql.DB) {
	api := g.Group("/api")
	{
		usersRoute(api, db)
	}
}

func usersRoute(api *gin.RouterGroup, db *sql.DB) {
	userController := controllers.NewUserController(db)
	api.GET("/users", userController.GetUsers)
	api.POST("/users", userController.AddNewUser)
}
