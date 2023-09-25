package main

import (
	"ginEssential/controller"
	"github.com/gin-gonic/gin"
)

func CollectRouters(r *gin.Engine) *gin.Engine {
	r.GET("/api/auth/register", controller.Register)

	return r
}
