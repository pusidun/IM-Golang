package server

import (
	"im-golang/controller"
	"im-golang/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", controller.Ping)
		v1.POST("user/login", controller.UserLogin)
		v1.POST("user/register", controller.UserRegister)
		v1.GET("user/info", middleware.AuthMiddleWare(), controller.Info)
	}

	return r
}
