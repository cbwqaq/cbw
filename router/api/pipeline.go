package api

import "github.com/gin-gonic/gin"

func CreateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"hello": "golang",
	})
}
