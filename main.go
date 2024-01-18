package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/user", func(ctx *gin.Context) {
		ctx.IndentedJSON(200, "Hello")
	})
	router.Run(":8080")
}
