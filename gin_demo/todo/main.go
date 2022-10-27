package main

import (
	todoRouter "go_demo/gin_demo/todo/router"
)

func main() {
	todoRouter.Run(":8081")
}
