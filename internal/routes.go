package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/sreerag-rajan/boilerplate-go/internal/templates"
)

// Routes object to hold all registered routes
type Routes struct {
	router *gin.Engine // Embed a Gin router instance
}

// NewRoutes creates a new Routes object with a Gin router
func NewRoutes(router *gin.Engine) *Routes {
	return &Routes{router: router}
}

// RegisterRoutes imports and registers routes from sub-packages
func (r *Routes) RegisterRoutes() {
	templates.RegisterRoutes(r.router)

}
