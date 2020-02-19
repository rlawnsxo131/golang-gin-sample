package main

import (
	"gin/api"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default() // create gin app
	api.ApplyRoutes(app)
	app.Run(":3001")
}
