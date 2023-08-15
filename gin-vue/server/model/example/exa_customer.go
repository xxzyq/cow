package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type ExaCustomer struct {
	global.GVA_MODEL
	CustomerName       string         `json:"customerName" form:"customerName" gorm:"comment:客户名"`                // 客户名
	CustomerPhoneData  string         `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:客户手机号"`    // 客户手机号
	SysUserID          uint           `json:"sysUserId" form:"sysUserId" gorm:"comment:管理ID"`                     // 管理ID
	SysUserAuthorityID uint           `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:管理角色ID"` // 管理角色ID
	SysUser            system.SysUser `json:"sysUser" form:"sysUser" gorm:"comment:管理详情"`                         // 管理详情
}

type ExaGoodType struct {
	global.GVA_MODEL
	GoodTypeId   int    `json:"goodTypeId" form:"goodTypeId" gorm:"comment:类型id"`    // 库存类别id
	GoodTypeName string `json:"goodTypeName" form:"goodTypeName" gorm:"comment:类型名"` // 库存名
	GoodTypeUnit string `json:"goodTypeUnit" form:"goodTypeUnit" gorm:"comment:类型名"` // 单位
	Operator     string `json:"operator" form:"operator" gorm:"comment:操作员"`         // 操作员

}
