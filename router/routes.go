package router

import (
	"github.com/gin-gonic/gin"
	"library/controller"
	"library/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("tokenAuth", middleware.TokenAuthHandler)

	r.POST("register", controller.SignUp)

	return r
}
