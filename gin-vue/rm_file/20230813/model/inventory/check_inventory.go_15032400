// 自动生成模板CheckInventory
package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// CheckInventory 结构体
type CheckInventory struct {
      global.GVA_MODEL
      GoodType  *int `json:"goodType" form:"goodType" gorm:"column:good_type;comment:库存种类;"`
      NewNumber  *int `json:"newNumber" form:"newNumber" gorm:"column:new_number;comment:数量;"`
      LossValue  *int `json:"lossValue" form:"lossValue" gorm:"column:loss_value;comment:损失值;"`
      Operator  string `json:"operator" form:"operator" gorm:"column:operator;comment:操作员;"`
}


// TableName CheckInventory 表名
func (CheckInventory) TableName() string {
  return "check_inventory"
}

