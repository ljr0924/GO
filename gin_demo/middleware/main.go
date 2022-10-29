package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Set("example", "12345")

		ctx.Next()

		latency := time.Since(t)
		log.Print(latency)
		// 获取发送的 status
		status := ctx.Writer.Status()
		log.Println(status)
	}
}

func Logger1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("这是logger1")
		ctx.Next()
	}
}

func Logger2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("这是logger2")
		ctx.Next()
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())
	r.Use(Logger1())
	r.Use(Logger2())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		// 打印："12345"
		log.Println(example)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
