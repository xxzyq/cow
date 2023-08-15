package weightRecord

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/weightRecord"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    weightRecordReq "github.com/flipped-aurora/gin-vue-admin/server/model/weightRecord/request"
)

type WeightRecordService struct {
}

// CreateWeightRecord 创建WeightRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (weightService *WeightRecordService) CreateWeightRecord(weight *weightRecord.WeightRecord) (err error) {
	err = global.GVA_DB.Create(weight).Error
	return err
}

// DeleteWeightRecord 删除WeightRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (weightService *WeightRecordService)DeleteWeightRecord(weight weightRecord.WeightRecord) (err error) {
	err = global.GVA_DB.Delete(&weight).Error
	return err
}

// DeleteWeightRecordByIds 批量删除WeightRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (weightService *WeightRecordService)DeleteWeightRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]weightRecord.WeightRecord{},"id in ?",ids.Ids).Error
	return err
}

// UpdateWeightRecord 更新WeightRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (weightService *WeightRecordService)UpdateWeightRecord(weight weightRecord.WeightRecord) (err error) {
	err = global.GVA_DB.Save(&weight).Error
	return err
}

// GetWeightRecord 根据id获取WeightRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (weightService *WeightRecordService)GetWeightRecord(id uint) (weight weightRecord.WeightRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&weight).Error
	return
}

// GetWeightRecordInfoList 分页获取WeightRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (weightService *WeightRecordService)GetWeightRecordInfoList(info weightRecordReq.WeightRecordSearch) (list []weightRecord.WeightRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&weightRecord.WeightRecord{})
    var weights []weightRecord.WeightRecord
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.RoughTime != "" {
        db = db.Where("rough_time LIKE ?","%"+ info.RoughTime+"%")
    }
    if info.NetTime != "" {
        db = db.Where("net_time LIKE ?","%"+ info.NetTime+"%")
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
         	orderMap["cowNumber"] = true
         	orderMap["routhWeight"] = true
         	orderMap["roughTime"] = true
         	orderMap["netWeight"] = true
         	orderMap["netTime"] = true
         	orderMap["operator"] = true
         	orderMap["changeNumber"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&weights).Error
	return  weights, total, err
}
