package route

import (
	createCategory "kyrol-marketplace-backend/controllers/category-controllers/create"
	deleteCategory "kyrol-marketplace-backend/controllers/category-controllers/delete"
	getCategory "kyrol-marketplace-backend/controllers/category-controllers/get"
	updateCategory "kyrol-marketplace-backend/controllers/category-controllers/update"
	middleware "kyrol-marketplace-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func InitCategoryRoute(route *gin.Engine) {
	group := route.Group("/api/v1/category")
	group.Use(middleware.Auth())
	group.GET("/:id", middleware.Gate("category_read"), getCategory.ShowCategory)
	group.GET("/", middleware.Gate("category_read"), getCategory.GetCategory)
	group.POST("/", middleware.Gate("category_create"), createCategory.CreateCategory)
	group.PUT("/:id", middleware.Gate("category_update"), updateCategory.UpdateCategory)
	group.PATCH("/:id", middleware.Gate("category_update"), updateCategory.UpdateCategory)
	group.DELETE("/:id", middleware.Gate("category_delete"), deleteCategory.DeleteCategory)
}
