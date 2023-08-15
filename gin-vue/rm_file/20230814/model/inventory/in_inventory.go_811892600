// 自动生成模板InInventory
package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	//"time"

)

// InInventory 结构体
type InInventory struct {
	global.GVA_MODEL
	GoodType *int   `json:"goodType" form:"goodType" gorm:"column:good_type;comment:库存类型;"`
	Number   *int   `json:"number" form:"number" gorm:"column:number;comment:入库数量;"`
	Unit     *int   `json:"unit" form:"unit" gorm:"column:unit;comment:库存单位;"`
	Operator string `json:"operator" form:"operator" gorm:"column:operator;comment:操作员;"`
}

// TableName InInventory 表名
func (InInventory) TableName() string {
	return "in_inventory"
}
