package routers

import (
	"server/controllers/api"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	apiRouters.Use(middlewares.InitMiddleware) //单个路由，单个过滤器
	{
		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.POST("/seterror", api.ApiController{}.SetError)
	}

}
