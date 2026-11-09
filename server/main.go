package main

import (
	"server/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//配置路由
	routers.ApiRoutersInit(r)
	r.Run(":8080") // 默认在 localhost:8080 启动服务

}
