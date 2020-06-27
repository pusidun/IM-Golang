package server

import (
	"im-golang/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", controller.Ping)
		v1.GET("user/login", controller.UserLogin)
		v1.POST("user/register", controller.UserRegister)
	}

	return r
}
