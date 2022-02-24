package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go_code/go-micro-todo/api-gateway/pkg/utils"
	service "go_code/go-micro-todo/api-gateway/services"
)

func UserRegister(ctx *gin.Context) {
	var userReq *service.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))
	userService := ctx.Keys["userService"].(service.UserService)
	userRep , err := userService.UserRegister(context.Background(), userReq)
	PanicIfUserError(err)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg": "注册成功",
		"data": userRep,
	})
}

func UserLogin(ctx *gin.Context) {
	var userReq *service.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))
	userService := ctx.Keys["userService"].(service.UserService)
	userRep , err := userService.UserLogin(context.Background(), userReq)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(uint(userRep.UserDetail.ID))
	PanicIfUserError(err)
	ctx.JSON(200, gin.H{
		"code": userRep.Code,
		"msg": "成功",
		"data": gin.H{
			"user": userRep.UserDetail,
			"token": token,
		},
	})
}
