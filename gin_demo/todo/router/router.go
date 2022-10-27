package router

import (
	"go_demo/gin_demo/todo/controller"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func InitRouter() {
	todoController := &controller.TodoController{}
	router.GET("/todo/list", todoController.GetList)
	router.GET("/todo/one", todoController.GetOne)
	router.POST("/todo/", todoController.Add)
	router.PUT("/todo/", todoController.Edit)
	router.DELETE("/todo/", todoController.Delete)
}

func Run(addr ...string) {
	InitRouter()
	router.Run(addr...)
}
