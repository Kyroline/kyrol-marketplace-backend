package route

import (
	createCategory "kyrol-marketplace-backend/controllers/category-controllers/create"
	deleteCategory "kyrol-marketplace-backend/controllers/category-controllers/delete"
	getCategory "kyrol-marketplace-backend/controllers/category-controllers/get"
	updateCategory "kyrol-marketplace-backend/controllers/category-controllers/update"

	"github.com/gin-gonic/gin"
)

func InitCategoryRoute(route *gin.Engine) {
	group := route.Group("/api/v1/category")
	group.GET("/:id", getCategory.ShowCategory)
	group.GET("/", getCategory.GetCategory)
	group.POST("/", createCategory.CreateCategory)
	group.PUT("/:id", updateCategory.UpdateCategory)
	group.PATCH("/:id", updateCategory.UpdateCategory)
	group.DELETE("/:id", deleteCategory.DeleteCategory)
}
