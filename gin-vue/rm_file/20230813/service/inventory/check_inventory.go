package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/inventory"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    inventoryReq "github.com/flipped-aurora/gin-vue-admin/server/model/inventory/request"
)

type CheckInventoryService struct {
}

// CreateCheckInventory 创建CheckInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkService *CheckInventoryService) CreateCheckInventory(check *inventory.CheckInventory) (err error) {
	err = global.GVA_DB.Create(check).Error
	return err
}

// DeleteCheckInventory 删除CheckInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkService *CheckInventoryService)DeleteCheckInventory(check inventory.CheckInventory) (err error) {
	err = global.GVA_DB.Delete(&check).Error
	return err
}

// DeleteCheckInventoryByIds 批量删除CheckInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkService *CheckInventoryService)DeleteCheckInventoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]inventory.CheckInventory{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCheckInventory 更新CheckInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkService *CheckInventoryService)UpdateCheckInventory(check inventory.CheckInventory) (err error) {
	err = global.GVA_DB.Save(&check).Error
	return err
}

// GetCheckInventory 根据id获取CheckInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkService *CheckInventoryService)GetCheckInventory(id uint) (check inventory.CheckInventory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&check).Error
	return
}

// GetCheckInventoryInfoList 分页获取CheckInventory记录
// Author [piexlmax](https://github.com/piexlmax)
func (checkService *CheckInventoryService)GetCheckInventoryInfoList(info inventoryReq.CheckInventorySearch) (list []inventory.CheckInventory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&inventory.CheckInventory{})
    var checks []inventory.CheckInventory
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.GoodType != nil {
        db = db.Where("good_type = ?",info.GoodType)
    }
    if info.LossValue != nil {
        db = db.Where("loss_value > ?",info.LossValue)
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
         	orderMap["goodType"] = true
         	orderMap["newNumber"] = true
         	orderMap["lossValue"] = true
         	orderMap["operator"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&checks).Error
	return  checks, total, err
}
