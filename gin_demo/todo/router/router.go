package router

import (
	todoController "go_demo/gin_demo/todo/controller"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func InitRouter() {
	todoRouter := router.Group("/todo")
	{
		todoRouter.GET("/list", todoController.GetList)
		todoRouter.GET("/one", todoController.GetOne)
		todoRouter.POST("/", todoController.Add)
		todoRouter.PUT("/", todoController.Edit)
		todoRouter.DELETE("/", todoController.Delete)
	}
}

func Run(addr ...string) {
	InitRouter()
	router.Run(addr...)
}
