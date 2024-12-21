package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nglLike/controllers"
)

func SetupRoutes(g *gin.Engine, db *sqlx.DB) {
	api := g.Group("/api")
	{
		usersRoute(api, db)
	}
}

func usersRoute(api *gin.RouterGroup, db *sqlx.DB) {
	userController := controllers.NewAuthController(db)
	api.POST("/auth/register", userController.RegisterHandler)
	api.POST("/auth/login", userController.LoginHandler)
}
