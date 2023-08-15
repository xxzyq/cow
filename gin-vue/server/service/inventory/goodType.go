package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type GoodTypeService struct{}

//func (m *GoodType) AddTypeS() {
//
//	m.AddGoodType()
//	fmt.Println("111")
//}
//func (m *GoodType) AddGoodType() {
//	// 连接数据库
//	db, err := sql.Open("mysql", "root:123457@tcp(localhost:3306)/gva")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	// 准备插入语句
//	stmt, err := db.Prepare("INSERT INTO inventory_type (Type_name, Operator) VALUES (?, ?)")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer stmt.Close()
//
//	// 执行插入操作
//	_, err = stmt.Exec(m.type_name, m.operator)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("数据添加成功！")
//}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExaCustomer
//@description: 创建库存类型
//@param: e model.ExaGoodType
//@return: err error

func (exa *GoodTypeService) CreateExaGoodType(e example.ExaGoodType) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除库存类型
//@param: e model.ExaGoodType
//@return: err error

func (exa *GoodTypeService) DeleteExaGoodType(e example.ExaGoodType) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaCustomer
//@description: 获取库存类型信息
//@param: id uint
//@return: customer model.ExaGoodType, err error

func (exa *GoodTypeService) GetExaGoodType(id uint) (goodType example.ExaGoodType, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&goodType).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *GoodTypeService) GetGoodTypeInfoList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&example.ExaCustomer{})
	var a system.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	if err != nil {
		return
	}
	var dataId []uint
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}
	var CustomerList []example.ExaCustomer
	err = db.Where("sys_user_authority_id in ?", dataId).Count(&total).Error
	if err != nil {
		return CustomerList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in ?", dataId).Find(&CustomerList).Error
	}
	return CustomerList, total, err
}
