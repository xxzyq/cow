package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InInventoryRouter struct {
}

// InitInInventoryRouter 初始化 InInventory 路由信息
func (s *InInventoryRouter) InitInInventoryRouter(Router *gin.RouterGroup) {
	InRouter := Router.Group("In").Use(middleware.OperationRecord())
	InRouterWithoutRecord := Router.Group("In")
	var InApi = v1.ApiGroupApp.InventoryApiGroup.InInventoryApi
	//var Api = v1.ApiGroupApp.InventoryApiGroup.InventoryDisplayApi
	{
		InRouter.POST("createInInventory", InApi.CreateInInventory)             // 新建InInventory
		InRouter.DELETE("deleteInInventory", InApi.DeleteInInventory)           // 删除InInventory
		InRouter.DELETE("deleteInInventoryByIds", InApi.DeleteInInventoryByIds) // 批量删除InInventory
		InRouter.PUT("updateInInventory", InApi.UpdateInInventory)              // 更新InInventory

	}
	{
		InRouterWithoutRecord.GET("findInInventory", InApi.FindInInventory)       // 根据ID获取InInventory
		InRouterWithoutRecord.GET("getInInventoryList", InApi.GetInInventoryList) // 获取InInventory列表
	}
}
