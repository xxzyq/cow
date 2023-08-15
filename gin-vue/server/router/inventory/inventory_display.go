package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InventoryDisplayRouter struct {
}

// InitInventoryDisplayRouter 初始化 InventoryDisplay 路由信息
func (s *InventoryDisplayRouter) InitInventoryDisplayRouter(Router *gin.RouterGroup) {
	inventDisRouter := Router.Group("inventDis").Use(middleware.OperationRecord())
	inventDisRouterWithoutRecord := Router.Group("inventDis")
	var inventDisApi = v1.ApiGroupApp.InventoryApiGroup.InventoryDisplayApi
	{
		inventDisRouter.POST("createInventoryDisplay", inventDisApi.CreateInventoryDisplay)             // 新建InventoryDisplay
		inventDisRouter.DELETE("deleteInventoryDisplay", inventDisApi.DeleteInventoryDisplay)           // 删除InventoryDisplay
		inventDisRouter.DELETE("deleteInventoryDisplayByIds", inventDisApi.DeleteInventoryDisplayByIds) // 批量删除InventoryDisplay
		inventDisRouter.DELETE("deleteInventoryDisplayByGoodType", inventDisApi.DeleteInventoryDisplayByGoodType)
		inventDisRouter.PUT("updateInventoryDisplay", inventDisApi.UpdateInventoryDisplay) // 更新InventoryDisplay
	}
	{
		inventDisRouterWithoutRecord.GET("findInventoryDisplay", inventDisApi.FindInventoryDisplay) // 根据ID获取InventoryDisplay
		inventDisRouterWithoutRecord.GET("findInventoryDisplayByGoodType", inventDisApi.FindInventoryDisplayByGoodType)
		inventDisRouterWithoutRecord.GET("getInventoryDisplayList", inventDisApi.GetInventoryDisplayList) // 获取InventoryDisplay列表
	}
}
