package router

import (
	v1 "GINCHAT/service/api/v1/pipeline"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	pipeline := r.Group("/api/v1")
	{
		pipeline.POST("/hello", v1.CreateUser)
	}
	return r
}
