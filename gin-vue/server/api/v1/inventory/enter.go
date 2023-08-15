package inventory

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	MyApi
	GoodTypeApi
	OutInventoryApi
	CheckInventoryApi
	InventoryDisplayApi
	InInventoryApi
}

var (
	myApiService	= service.ServiceGroupApp.InventoryServiceGroup.MyApi
	goodTypeService	= service.ServiceGroupApp.InventoryServiceGroup.GoodTypeService
)
