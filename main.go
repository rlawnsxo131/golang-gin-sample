package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rlawnsxo131/golang-gin-sample/api"
)

func main() {
	app := gin.Default() // create gin app
	api.ApplyRouters(app)
	app.Run(":3001")
}
