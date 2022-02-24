package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go_code/go-micro-todo/api-gateway/pkg/utils"
	service "go_code/go-micro-todo/api-gateway/services"
	"strconv"
)

func GetTaskList(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	taskService := ctx.Keys["taskService"].(service.TaskService)
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	//调用服务
	res, err := taskService.GetTasksList(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg": "获取备忘录列表成功",
		"data": gin.H{
			"list": res.TaskList,
			"count": res.Count,
		},
	})
}

func CreateTask(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	taskService := ctx.Keys["taskService"].(service.TaskService)
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	res, err := taskService.CreateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg": "创建备忘录成功",
		"data": res.TaskDetail,
	})
}

func GetTask(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	taskService := ctx.Keys["taskService"].(service.TaskService)
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(ctx.Param("id"))
	taskReq.Id = uint64(id)
	res, err := taskService.GetTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg": "获取备忘录详情成功",
		"data": res.TaskDetail,
	})
}

func UpdateTask(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	taskService := ctx.Keys["taskService"].(service.TaskService)
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(ctx.Param("id"))
	taskReq.Id = uint64(id)
	res, err := taskService.UpdateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg": "更新备忘录成功",
		"data": res.TaskDetail,
	})
}

func DeleteTask(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))
	taskService := ctx.Keys["taskService"].(service.TaskService)
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(ctx.Param("id"))
	taskReq.Id = uint64(id)
	res, err := taskService.DeleteTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg": "删除备忘录成功",
		"data": res.TaskDetail,
	})
}
