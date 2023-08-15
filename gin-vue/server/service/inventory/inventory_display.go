package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/inventory"
	inventoryReq "github.com/flipped-aurora/gin-vue-admin/server/model/inventory/request"
)

type InventoryDisplayService struct {
}

// CreateInventoryDisplay 创建InventoryDisplay记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventDisService *InventoryDisplayService) CreateInventoryDisplay(inventDis *inventory.InventoryDisplay) (err error) {
	err = global.GVA_DB.Create(inventDis).Error
	return err
}

// DeleteInventoryDisplay 删除InventoryDisplay记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventDisService *InventoryDisplayService) DeleteInventoryDisplay(inventDis inventory.InventoryDisplay) (err error) {
	err = global.GVA_DB.Delete(&inventDis).Error
	return err
}

// DeleteInventoryDisplayByIds 批量删除InventoryDisplay记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventDisService *InventoryDisplayService) DeleteInventoryDisplayByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]inventory.InventoryDisplay{}, "id in ?", ids.Ids).Error
	return err
}

// DeleteInventoryDisplayByGoodType 根据goodType删除InventoryDisplay记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventDisService *InventoryDisplayService) DeleteInventoryDisplayByGoodType(goodType request.GoodTypeReq) (err error) {
	err = global.GVA_DB.Delete(&[]inventory.InventoryDisplay{}, "good_type in ?", goodType.GoodType).Error
	return err
}

// UpdateInventoryDisplay 更新InventoryDisplay记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventDisService *InventoryDisplayService) UpdateInventoryDisplay(inventDis inventory.InventoryDisplay) (err error) {
	err = global.GVA_DB.Save(&inventDis).Error
	return err
}

//// 根据传入的结构体进行查询，如果存在则根据 good_type 进行更新
//func (inventDisService *InventoryDisplayService) UpdateInventoryDisplayByGoodType(inventDis inventory.InventoryDisplay) error {
//	var existingRecord models.YourModelType
//	err := inventDisService.DB.Where("good_type = ?", inventDis.GoodType).First(&existingRecord).Error
//	if err != nil {
//		if gorm.IsRecordNotFoundError(err) {
//			// 不存在记录，执行创建操作
//			return inventDisService.DB.Create(inventDis).Error
//		}
//		return err
//	}
//
//	// 存在记录，根据 good_type 进行更新
//	return inventDisService.DB.Model(&existingRecord).Updates(inventDis).Error
//}

// GetInventoryDisplay 根据id获取InventoryDisplay记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventDisService *InventoryDisplayService) GetInventoryDisplay(id uint) (inventDis inventory.InventoryDisplay, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&inventDis).Error
	return
}

// GetInventoryDisplay 根据goodType获取InventoryDisplay记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventDisService *InventoryDisplayService) GetInventoryDisplayByGoodType(goodType int) (inventDis inventory.InventoryDisplay, err error) {
	err = global.GVA_DB.Where("good_type = ?", goodType).First(&inventDis).Error
	return
}

// GetInventoryDisplayInfoList 分页获取InventoryDisplay记录
// Author [piexlmax](https://github.com/piexlmax)
func (inventDisService *InventoryDisplayService) GetInventoryDisplayInfoList(info inventoryReq.InventoryDisplaySearch) (list []inventory.InventoryDisplay, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&inventory.InventoryDisplay{})
	var inventDiss []inventory.InventoryDisplay
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["goodType"] = true
	orderMap["number"] = true
	orderMap["unit"] = true
	orderMap["totalLossValue"] = true
	orderMap["totalValue"] = true
	orderMap["totalLossRate"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&inventDiss).Error
	return inventDiss, total, err
}
