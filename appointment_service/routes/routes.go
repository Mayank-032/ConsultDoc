package routes

import (
	"healthcare-service/bootstrap"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api");
	bootstrap.Init(apiGroup)
}