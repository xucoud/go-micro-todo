package weblib

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_code/go-micro-todo/api-gateway/weblib/handlers"
	"go_code/go-micro-todo/api-gateway/weblib/middleware"

	"github.com/gin-contrib/sessions/cookie"
)

func NewRouter(services ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(services), middleware.ErrorMiddleware())

	store := cookie.NewStore([]byte("secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))

	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
				ctx.JSON(200, "success")
		})
		v1.POST("/user/register", handlers.UserRegister)
		v1.POST("/user/login", handlers.UserLogin)

		//备忘录
		auth := v1.Group("/")
		auth.Use(middleware.JWT())
		{
			auth.GET("/tasks", handlers.GetTaskList)
			auth.GET("/task/:id", handlers.GetTask)
			auth.POST("/task", handlers.CreateTask)
			auth.PUT("/task/:id", handlers.UpdateTask)
			auth.DELETE("/task/:id", handlers.DeleteTask)
		}

	}

	return ginRouter
}
