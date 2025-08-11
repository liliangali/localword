package router

import (
	"github.com/gin-gonic/gin"
	"localword/order-web/api/order"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("word").Use()
	{
		//GoodsRouter.GET("", goods.List) //商品列表
		//GoodsRouter.POST("assign", order.Assign) //改接口需要管理员权限
		GoodsRouter.POST("wordExtend", order.WordBaidu)
		GoodsRouter.POST("wordId", order.ExtendWordById)

	}
}
