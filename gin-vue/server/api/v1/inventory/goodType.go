package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	exampleRes "github.com/flipped-aurora/gin-vue-admin/server/model/example/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GoodTypeApi struct{}

// CreateExaGoodType
// @Tags      ExaGoodType
// @Summary   创建库存类型
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaGoodType            true  "库存类型, 库存单位"
// @Success   200   {object}  response.Response{msg=string}  "创建库存类型"
// @Router    /goodType/goodType [post]
func (e *GoodTypeApi) CreateExaGoodType(c *gin.Context) {
	var goodType example.ExaGoodType
	err := c.ShouldBindJSON(&goodType) //检查json
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(goodType, utils.GoodTypeVerify) //检查所有数据类型
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//goodType.GoodTypeId = utils.GetUserID(c)
	err = goodTypeService.CreateExaGoodType(goodType) //执行数据库操作
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteExaGoodType
// @Tags      ExaGoodType
// @Summary   删除库存类型
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaGoodType            true  "库存类型ID"
// @Success   200   {object}  response.Response{msg=string}  "删除库存类型"
// @Router    /goodType/goodType [delete]
func (e *GoodTypeApi) DeleteExaGoodType(c *gin.Context) {
	var goodType example.ExaGoodType
	err := c.ShouldBindJSON(&goodType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(goodType.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = goodTypeService.DeleteExaGoodType(goodType)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetExaGoodType
// @Tags      ExaGoodType
// @Summary   获取单一商品类别信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     example.ExaGoodType                                                true  "库存类型ID"
// @Success   200   {object}  response.Response{data=exampleRes.ExaGoodTypeResponse,msg=string}  "获取单一库存类别信息,返回包括库存类别详情"
// @Router    /goodType/goodType [get]
func (e *GoodTypeApi) GetExaGoodType(c *gin.Context) {
	var goodType example.ExaGoodType
	err := c.ShouldBindQuery(&goodType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(goodType.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := goodTypeService.GetExaGoodType(goodType.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(exampleRes.ExaGoodTypeResponse{GoodType: data}, "获取成功", c)
}

// GetExaGoodTypeList
// @Tags      ExaGoodType
// @Summary   分页获取权限客户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /goodType/goodTypeList [get]
func (e *GoodTypeApi) GetExaGoodTypeList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	goodTypeList, total, err := goodTypeService.GetGoodTypeInfoList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     goodTypeList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
