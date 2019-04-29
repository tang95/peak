package routers

import (
	"github.com/gin-gonic/gin"
	"peak/controllers"
	"peak/handlers"
)

func RegisterRoute(router *gin.Engine) {
	//主路由
	ctrl := new(controllers.MainController)
	router.GET("/", ctrl.Main)

	//项目路由
	project := router.Group("/api/project")
	{
		ctrl := new(controllers.ProjectController)
		project.POST("/", ctrl.Create)
		project.GET("/", ctrl.List)
	}

	//socketio路由
	router.GET("/socket.io/", handlers.SocketIOHandler())
	router.POST("/socket.io/", handlers.SocketIOHandler())
	router.Handle("WS", "/socket.io", handlers.SocketIOHandler())
	router.Handle("WSS", "/socket.io", handlers.SocketIOHandler())
}
