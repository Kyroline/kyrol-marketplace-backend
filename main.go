package main

import (
	database "kyrol-marketplace-backend/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	route "kyrol-marketplace-backend/routes"
)

func main() {
	database.Conn()
	database.Migrate()

	router := SetupRouter()
	router.Run()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	route.InitProductRoute(router)
	route.InitAuthRoute(router)

	return router
}
