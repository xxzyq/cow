// 自动生成模板InventoryDisplay
package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// InventoryDisplay 结构体
type InventoryDisplay struct {
	global.GVA_MODEL
	GoodType       int      `json:"goodType" form:"goodType" gorm:"column:good_type;comment:库存类型;"`
	Number         *int     `json:"number" form:"number" gorm:"column:number;comment:数量;"`
	Unit           *int     `json:"unit" form:"unit" gorm:"column:unit;comment:库存单位;"`
	TotalLossValue *int     `json:"totalLossValue" form:"totalLossValue" gorm:"column:total_loss_value;comment:总损失值;"`
	TotalValue     *int     `json:"totalValue" form:"totalValue" gorm:"column:total_value;comment:历史累计值;"`
	TotalLossRate  *float64 `json:"totalLossRate" form:"totalLossRate" gorm:"column:total_loss_rate;comment:总损失数/总数;"`
}

// TableName InventoryDisplay 表名
func (InventoryDisplay) TableName() string {
	return "inventory"
}
