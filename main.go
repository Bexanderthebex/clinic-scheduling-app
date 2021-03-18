package main

import (
	gin "github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.Group("/appointments")

	route.Run(":5000")
}
