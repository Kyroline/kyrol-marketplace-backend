package main

import (
	database "kyrol-marketplace-backend/database"

	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	route "kyrol-marketplace-backend/routes"
)

func main() {
	database.Conn()
	database.Migrate()
	if err := database.Seeder(); err != nil {
		fmt.Printf("Error : ", err.Error())
		return
	}

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
	route.InitCategoryRoute(router)
	route.InitCartRoute(router)
	route.InitInvoiceRoute(router)

	return router
}
