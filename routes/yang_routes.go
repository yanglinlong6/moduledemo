package routes

import (
	"github.com/gin-gonic/gin"
	"moduledemo/controller"
)

func InitRoutes() *gin.Engine {
	gin.SetMode("debug")

	r := gin.Default()
	//r.Use(gin.Recovery())

	apiGroup := r.Group("/")

	InitBaseRoutes(apiGroup)

	return r
}

func InitBaseRoutes(r *gin.RouterGroup) gin.IRoutes {
	base := r.Group("/yang")
	{
		base.GET("yang", controller.Yang.Hello)
		base.GET("save", controller.Yang.Save)
		base.GET("ping", controller.Demo)
	}

	return r
}
