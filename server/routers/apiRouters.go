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
		apiRouters.GET("/getDeposit", api.ApiController{}.GetDeposit)
		apiRouters.GET("/getCancel", api.ApiController{}.GetCancel)
		apiRouters.GET("/getOrder", api.ApiController{}.GetOrder)
		apiRouters.GET("/getTrade", api.ApiController{}.GetTrade)
		apiRouters.GET("/getWithdraw", api.ApiController{}.GetWithdraw)
		apiRouters.GET("/getApproval", api.ApiController{}.GetApproval)
		apiRouters.GET("/getTransfer", api.ApiController{}.GetTransfer)
	}

}
