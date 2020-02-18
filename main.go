package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "world",
			})
		})
	}

	router.Run(":3001")
}
