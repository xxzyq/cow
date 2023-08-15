package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OutInventoryRouter struct {
}

// InitOutInventoryRouter 初始化 OutInventory 路由信息
func (s *OutInventoryRouter) InitOutInventoryRouter(Router *gin.RouterGroup) {
	outRouter := Router.Group("out").Use(middleware.OperationRecord())
	outRouterWithoutRecord := Router.Group("out")
	var outApi = v1.ApiGroupApp.InventoryApiGroup.OutInventoryApi
	{
		outRouter.POST("createOutInventory", outApi.CreateOutInventory)   // 新建OutInventory
		outRouter.DELETE("deleteOutInventory", outApi.DeleteOutInventory) // 删除OutInventory
		outRouter.DELETE("deleteOutInventoryByIds", outApi.DeleteOutInventoryByIds) // 批量删除OutInventory
		outRouter.PUT("updateOutInventory", outApi.UpdateOutInventory)    // 更新OutInventory
	}
	{
		outRouterWithoutRecord.GET("findOutInventory", outApi.FindOutInventory)        // 根据ID获取OutInventory
		outRouterWithoutRecord.GET("getOutInventoryList", outApi.GetOutInventoryList)  // 获取OutInventory列表
	}
}
