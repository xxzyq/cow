package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/inventory"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/weightRecord"
)

type ServiceGroup struct {
	SystemServiceGroup       system.ServiceGroup
	ExampleServiceGroup      example.ServiceGroup
	InventoryServiceGroup    inventory.ServiceGroup
	WeightRecordServiceGroup weightRecord.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
