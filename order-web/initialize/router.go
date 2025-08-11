package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"localword/order-web/middlewares"
	"localword/order-web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("/o/v1")

	zap.S().Infof("启动用户信息")
	router.InitOrderRouter(ApiGroup)

	return Router
}
