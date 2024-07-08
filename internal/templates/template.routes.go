package templates

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes defines routes for feature1 and registers them with the provided router
func RegisterRoutes(router *gin.Engine) {
	template := router.Group("/api/v1/template") // Define a template for feature1 routes
	{
		template.GET("/", Gettemplates)
		template.POST("/", Createtemplates)
		template.GET("/:id", Gettemplate)
		template.DELETE("/:id", Deletetemplate)
		template.PATCH("/:id", Updatetemplate)
	}
}
