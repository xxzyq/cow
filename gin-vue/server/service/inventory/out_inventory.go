package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/inventory"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    inventoryReq "github.com/flipped-aurora/gin-vue-admin/server/model/inventory/request"
)

type OutInventoryService struct {
}

// CreateOutInventory 创建OutInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (outService *OutInventoryService) CreateOutInventory(out *inventory.OutInventory) (err error) {
	err = global.GVA_DB.Create(out).Error
	return err
}

// DeleteOutInventory 删除OutInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (outService *OutInventoryService)DeleteOutInventory(out inventory.OutInventory) (err error) {
	err = global.GVA_DB.Delete(&out).Error
	return err
}

// DeleteOutInventoryByIds 批量删除OutInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (outService *OutInventoryService)DeleteOutInventoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]inventory.OutInventory{},"id in ?",ids.Ids).Error
	return err
}

// UpdateOutInventory 更新OutInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (outService *OutInventoryService)UpdateOutInventory(out inventory.OutInventory) (err error) {
	err = global.GVA_DB.Save(&out).Error
	return err
}

// GetOutInventory 根据id获取OutInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (outService *OutInventoryService)GetOutInventory(id uint) (out inventory.OutInventory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&out).Error
	return
}

// GetOutInventoryInfoList 分页获取OutInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (outService *OutInventoryService)GetOutInventoryInfoList(info inventoryReq.OutInventorySearch) (list []inventory.OutInventory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&inventory.OutInventory{})
    var outs []inventory.OutInventory
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.TypeName != nil {
        db = db.Where("type_name = ?",info.TypeName)
    }
    if info.Operator != "" {
        db = db.Where("operator LIKE ?","%"+ info.Operator+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["typeName"] = true
         	orderMap["number"] = true
         	orderMap["operator"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&outs).Error
	return  outs, total, err
}
