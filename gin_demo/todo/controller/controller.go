package controller

import "github.com/gin-gonic/gin"

type Todo struct {
	Content string `json:"content"`
}

type TodoController struct {
}

func (tc *TodoController) GetList(c *gin.Context) {
	todoList := []*Todo{
		{"todo1"},
		{"todo2"},
		{"todo3"},
	}
	c.JSON(200, gin.H{
		"code": 100000,
		"msg":  "success",
		"data": todoList,
	})
}

func (tc *TodoController) GetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 100000,
		"msg":  "success",
		"data": &Todo{},
	})
}

func (tc *TodoController) Add(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 100000,
		"msg":  "success",
	})
}

func (tc *TodoController) Edit(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 100000,
		"msg":  "success",
	})
}

func (tc *TodoController) Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 100000,
		"msg":  "success",
	})
}
