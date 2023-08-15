package inventory

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type GoodTypeRouter struct{}

func (m *GoodTypeRouter) InitTypeRouter(Router *gin.RouterGroup) {
	goodTypeRouter := Router.Group("goodType")
	exaGoodTypeApi := v1.ApiGroupApp.InventoryApiGroup.GoodTypeApi
	{
		goodTypeRouter.POST("goodType", exaGoodTypeApi.CreateExaGoodType)
		goodTypeRouter.DELETE("goodType", exaGoodTypeApi.DeleteExaGoodType)
	}
	{
		goodTypeRouter.GET("goodType", exaGoodTypeApi.GetExaGoodType)
		goodTypeRouter.GET("goodTypeList", exaGoodTypeApi.GetExaGoodTypeList)
	}

}
