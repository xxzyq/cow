package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/inventory"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/weightRecord"
)

type RouterGroup struct {
	System       system.RouterGroup
	Example      example.RouterGroup
	Inventory    inventory.RouterGroup
	WeightRecord weightRecord.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
