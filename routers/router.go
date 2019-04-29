package routers

import (
	"github.com/gin-gonic/gin"
	"peak/controllers"
)

func RegisterRoute(router *gin.Engine) {
	project := router.Group("/api/project")
	{
		ctrl := new(controllers.ProjectController)
		project.POST("/", ctrl.Create)
	}
}
