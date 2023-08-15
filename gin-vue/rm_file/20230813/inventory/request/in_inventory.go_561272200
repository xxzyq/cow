package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/inventory"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type InInventorySearch struct{
    inventory.InInventory
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartDate  *time.Time  `json:"startDate" form:"startDate"`
    EndDate  *time.Time  `json:"endDate" form:"endDate"`
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
