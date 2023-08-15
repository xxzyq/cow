package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CheckInventoryRouter struct {
}

// InitCheckInventoryRouter 初始化 CheckInventory 路由信息
func (s *CheckInventoryRouter) InitCheckInventoryRouter(Router *gin.RouterGroup) {
	checkRouter := Router.Group("check").Use(middleware.OperationRecord())
	checkRouterWithoutRecord := Router.Group("check")
	var checkApi = v1.ApiGroupApp.InventoryApiGroup.CheckInventoryApi
	{
		checkRouter.POST("createCheckInventory", checkApi.CreateCheckInventory)   // 新建CheckInventory
		checkRouter.DELETE("deleteCheckInventory", checkApi.DeleteCheckInventory) // 删除CheckInventory
		checkRouter.DELETE("deleteCheckInventoryByIds", checkApi.DeleteCheckInventoryByIds) // 批量删除CheckInventory
		checkRouter.PUT("updateCheckInventory", checkApi.UpdateCheckInventory)    // 更新CheckInventory
	}
	{
		checkRouterWithoutRecord.GET("findCheckInventory", checkApi.FindCheckInventory)        // 根据ID获取CheckInventory
		checkRouterWithoutRecord.GET("getCheckInventoryList", checkApi.GetCheckInventoryList)  // 获取CheckInventory列表
	}
}
