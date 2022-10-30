package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Content string `json:"content"`
}

func GetList(c *gin.Context) {
	var args GetListArgs
	err := c.ShouldBindQuery(&args)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, ArgsErrResp(err.Error()))
		return
	}

	log.Println(args)

	todoList := []*Todo{
		{"todo1"},
		{"todo2"},
		{"todo3"},
	}
	c.JSON(200, SuccessRespWithData(todoList))
}

func GetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 100000,
		"msg":  "success",
		"data": &Todo{},
	})
}

func Add(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 100000,
		"msg":  "success",
	})
}

func Edit(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 100000,
		"msg":  "success",
	})
}

func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 100000,
		"msg":  "success",
	})
}
