// 自动生成模板OutInventory
package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// OutInventory 结构体
type OutInventory struct {
      global.GVA_MODEL
      TypeName  *int `json:"typeName" form:"typeName" gorm:"column:type_name;comment:库存类型;"`
      Number  *int `json:"number" form:"number" gorm:"column:number;comment:数量;"`
      Operator  string `json:"operator" form:"operator" gorm:"column:operator;comment:操作员;"`
}


// TableName OutInventory 表名
func (OutInventory) TableName() string {
  return "out_inventory"
}

