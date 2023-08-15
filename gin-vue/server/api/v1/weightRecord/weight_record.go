package weightRecord

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/weightRecord"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    weightRecordReq "github.com/flipped-aurora/gin-vue-admin/server/model/weightRecord/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type WeightRecordApi struct {
}

var weightService = service.ServiceGroupApp.WeightRecordServiceGroup.WeightRecordService


// CreateWeightRecord 创建WeightRecord
// @Tags WeightRecord
// @Summary 创建WeightRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weightRecord.WeightRecord true "创建WeightRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weight/createWeightRecord [post]
func (weightApi *WeightRecordApi) CreateWeightRecord(c *gin.Context) {
	var weight weightRecord.WeightRecord
	err := c.ShouldBindJSON(&weight)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "CowNumber":{utils.NotEmpty()},
        "RouthWeight":{utils.NotEmpty()},
        "NetWeight":{utils.NotEmpty()},
    }
	if err := utils.Verify(weight, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := weightService.CreateWeightRecord(&weight); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWeightRecord 删除WeightRecord
// @Tags WeightRecord
// @Summary 删除WeightRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weightRecord.WeightRecord true "删除WeightRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weight/deleteWeightRecord [delete]
func (weightApi *WeightRecordApi) DeleteWeightRecord(c *gin.Context) {
	var weight weightRecord.WeightRecord
	err := c.ShouldBindJSON(&weight)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := weightService.DeleteWeightRecord(weight); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWeightRecordByIds 批量删除WeightRecord
// @Tags WeightRecord
// @Summary 批量删除WeightRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WeightRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /weight/deleteWeightRecordByIds [delete]
func (weightApi *WeightRecordApi) DeleteWeightRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := weightService.DeleteWeightRecordByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWeightRecord 更新WeightRecord
// @Tags WeightRecord
// @Summary 更新WeightRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body weightRecord.WeightRecord true "更新WeightRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /weight/updateWeightRecord [put]
func (weightApi *WeightRecordApi) UpdateWeightRecord(c *gin.Context) {
	var weight weightRecord.WeightRecord
	err := c.ShouldBindJSON(&weight)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "CowNumber":{utils.NotEmpty()},
          "RouthWeight":{utils.NotEmpty()},
          "NetWeight":{utils.NotEmpty()},
      }
    if err := utils.Verify(weight, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := weightService.UpdateWeightRecord(weight); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWeightRecord 用id查询WeightRecord
// @Tags WeightRecord
// @Summary 用id查询WeightRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weightRecord.WeightRecord true "用id查询WeightRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /weight/findWeightRecord [get]
func (weightApi *WeightRecordApi) FindWeightRecord(c *gin.Context) {
	var weight weightRecord.WeightRecord
	err := c.ShouldBindQuery(&weight)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reweight, err := weightService.GetWeightRecord(weight.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reweight": reweight}, c)
	}
}

// GetWeightRecordList 分页获取WeightRecord列表
// @Tags WeightRecord
// @Summary 分页获取WeightRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query weightRecordReq.WeightRecordSearch true "分页获取WeightRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weight/getWeightRecordList [get]
func (weightApi *WeightRecordApi) GetWeightRecordList(c *gin.Context) {
	var pageInfo weightRecordReq.WeightRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := weightService.GetWeightRecordInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
