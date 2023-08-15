// 自动生成模板WeightRecord
package weightRecord

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"

	"gorm.io/datatypes"
)

// WeightRecord 结构体
type WeightRecord struct {
	global.GVA_MODEL
	CowNumber    *int           `json:"cowNumber" form:"cowNumber" gorm:"column:cow_number;comment:称重顺序;"`
	RouthWeight  *int           `json:"routhWeight" form:"routhWeight" gorm:"column:routh_weight;comment:毛重;"`
	RoughTime    string         `json:"roughTime" form:"roughTime" gorm:"column:rough_time;comment:毛重时间;"`
	NetWeight    *int           `json:"netWeight" form:"netWeight" gorm:"column:net_weight;comment:酮体重;"`
	NetTime      string         `json:"netTime" form:"netTime" gorm:"column:net_time;comment:酮体称重时间;"`
	Operator     string         `json:"operator" form:"operator" gorm:"column:operator;comment:调整数据才会生成;"`
	ChangeNumber *int           `json:"changeNumber" form:"changeNumber" gorm:"column:change_number;comment:毛重的调整值;"`
	Remark       string         `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
	Photo1       string         `json:"photo1" form:"photo1" gorm:"column:photo1;comment:图片1;"`
	Photo2       datatypes.JSON `json:"photo2" form:"photo2" gorm:"column:photo2;comment:图片2;"`
	Photo3       datatypes.JSON `json:"photo3" form:"photo3" gorm:"column:photo3;comment:图片组3;"`
}

// TableName WeightRecord 表名
func (WeightRecord) TableName() string {
	return "weight_record"
}
