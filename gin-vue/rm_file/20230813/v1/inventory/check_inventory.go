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

type CheckInventoryApi struct {
}

var checkService = service.ServiceGroupApp.InventoryServiceGroup.CheckInventoryService


// CreateCheckInventory 创建CheckInventory
// @Tags CheckInventory
// @Summary 创建CheckInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventory.CheckInventory true "创建CheckInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /check/createCheckInventory [post]
func (checkApi *CheckInventoryApi) CreateCheckInventory(c *gin.Context) {
	var check inventory.CheckInventory
	err := c.ShouldBindJSON(&check)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "GoodType":{utils.NotEmpty()},
        "NewNumber":{utils.NotEmpty()},
        "Operator":{utils.NotEmpty()},
    }
	if err := utils.Verify(check, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := checkService.CreateCheckInventory(&check); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCheckInventory 删除CheckInventory
// @Tags CheckInventory
// @Summary 删除CheckInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventory.CheckInventory true "删除CheckInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /check/deleteCheckInventory [delete]
func (checkApi *CheckInventoryApi) DeleteCheckInventory(c *gin.Context) {
	var check inventory.CheckInventory
	err := c.ShouldBindJSON(&check)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := checkService.DeleteCheckInventory(check); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCheckInventoryByIds 批量删除CheckInventory
// @Tags CheckInventory
// @Summary 批量删除CheckInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CheckInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /check/deleteCheckInventoryByIds [delete]
func (checkApi *CheckInventoryApi) DeleteCheckInventoryByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := checkService.DeleteCheckInventoryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCheckInventory 更新CheckInventory
// @Tags CheckInventory
// @Summary 更新CheckInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventory.CheckInventory true "更新CheckInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /check/updateCheckInventory [put]
func (checkApi *CheckInventoryApi) UpdateCheckInventory(c *gin.Context) {
	var check inventory.CheckInventory
	err := c.ShouldBindJSON(&check)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "GoodType":{utils.NotEmpty()},
          "NewNumber":{utils.NotEmpty()},
          "Operator":{utils.NotEmpty()},
      }
    if err := utils.Verify(check, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := checkService.UpdateCheckInventory(check); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCheckInventory 用id查询CheckInventory
// @Tags CheckInventory
// @Summary 用id查询CheckInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query inventory.CheckInventory true "用id查询CheckInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /check/findCheckInventory [get]
func (checkApi *CheckInventoryApi) FindCheckInventory(c *gin.Context) {
	var check inventory.CheckInventory
	err := c.ShouldBindQuery(&check)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recheck, err := checkService.GetCheckInventory(check.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recheck": recheck}, c)
	}
}

// GetCheckInventoryList 分页获取CheckInventory列表
// @Tags CheckInventory
// @Summary 分页获取CheckInventory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query inventoryReq.CheckInventorySearch true "分页获取CheckInventory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /check/getCheckInventoryList [get]
func (checkApi *CheckInventoryApi) GetCheckInventoryList(c *gin.Context) {
	var pageInfo inventoryReq.CheckInventorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := checkService.GetCheckInventoryInfoList(pageInfo); err != nil {
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
