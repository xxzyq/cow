package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/inventory"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    inventoryReq "github.com/flipped-aurora/gin-vue-admin/server/model/inventory/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type OutInventoryApi struct {
}

var outService = service.ServiceGroupApp.InventoryServiceGroup.OutInventoryService


// CreateOutInventory 创建OutInventory
// @Tags OutInventory
// @Summary 创建OutInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventory.OutInventory true "创建OutInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /out/createOutInventory [post]
func (outApi *OutInventoryApi) CreateOutInventory(c *gin.Context) {
	var out inventory.OutInventory
	err := c.ShouldBindJSON(&out)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "TypeName":{utils.NotEmpty()},
        "Number":{utils.NotEmpty()},
        "Operator":{utils.NotEmpty()},
    }
	if err := utils.Verify(out, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := outService.CreateOutInventory(&out); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOutInventory 删除OutInventory
// @Tags OutInventory
// @Summary 删除OutInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventory.OutInventory true "删除OutInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /out/deleteOutInventory [delete]
func (outApi *OutInventoryApi) DeleteOutInventory(c *gin.Context) {
	var out inventory.OutInventory
	err := c.ShouldBindJSON(&out)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := outService.DeleteOutInventory(out); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOutInventoryByIds 批量删除OutInventory
// @Tags OutInventory
// @Summary 批量删除OutInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OutInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /out/deleteOutInventoryByIds [delete]
func (outApi *OutInventoryApi) DeleteOutInventoryByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := outService.DeleteOutInventoryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOutInventory 更新OutInventory
// @Tags OutInventory
// @Summary 更新OutInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventory.OutInventory true "更新OutInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /out/updateOutInventory [put]
func (outApi *OutInventoryApi) UpdateOutInventory(c *gin.Context) {
	var out inventory.OutInventory
	err := c.ShouldBindJSON(&out)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "TypeName":{utils.NotEmpty()},
          "Number":{utils.NotEmpty()},
          "Operator":{utils.NotEmpty()},
      }
    if err := utils.Verify(out, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := outService.UpdateOutInventory(out); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOutInventory 用id查询OutInventory
// @Tags OutInventory
// @Summary 用id查询OutInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query inventory.OutInventory true "用id查询OutInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /out/findOutInventory [get]
func (outApi *OutInventoryApi) FindOutInventory(c *gin.Context) {
	var out inventory.OutInventory
	err := c.ShouldBindQuery(&out)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reout, err := outService.GetOutInventory(out.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reout": reout}, c)
	}
}

// GetOutInventoryList 分页获取OutInventory列表
// @Tags OutInventory
// @Summary 分页获取OutInventory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query inventoryReq.OutInventorySearch true "分页获取OutInventory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /out/getOutInventoryList [get]
func (outApi *OutInventoryApi) GetOutInventoryList(c *gin.Context) {
	var pageInfo inventoryReq.OutInventorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := outService.GetOutInventoryInfoList(pageInfo); err != nil {
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
