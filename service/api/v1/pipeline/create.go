package pipeline

import (
	"GINCHAT/dao/api/v1"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var u v1.Student
	ctx.BindJSON(&u)
	v1.Create(u)
}
