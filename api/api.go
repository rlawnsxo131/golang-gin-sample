package api

import (
	"github.com/gin-gonic/gin"
)

func ApplyRouters(r *gin.Engine) {
	api := r.Group("/api")
	{
		apiv1.ApplyRouters(api)
	}
}
