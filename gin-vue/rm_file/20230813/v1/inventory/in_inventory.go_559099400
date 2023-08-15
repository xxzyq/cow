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

type InInventoryApi struct {
}

var InService = service.ServiceGroupApp.InventoryServiceGroup.InInventoryService


// CreateInInventory 创建InInventory
// @Tags InInventory
// @Summary 创建InInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventory.InInventory true "创建InInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /In/createInInventory [post]
func (InApi *InInventoryApi) CreateInInventory(c *gin.Context) {
	var In inventory.InInventory
	err := c.ShouldBindJSON(&In)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Date":{utils.NotEmpty()},
        "GoodType":{utils.NotEmpty()},
        "Number":{utils.NotEmpty()},
        "Unit":{utils.NotEmpty()},
        "Operator":{utils.NotEmpty()},
    }
	if err := utils.Verify(In, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := InService.CreateInInventory(&In); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteInInventory 删除InInventory
// @Tags InInventory
// @Summary 删除InInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventory.InInventory true "删除InInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /In/deleteInInventory [delete]
func (InApi *InInventoryApi) DeleteInInventory(c *gin.Context) {
	var In inventory.InInventory
	err := c.ShouldBindJSON(&In)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := InService.DeleteInInventory(In); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteInInventoryByIds 批量删除InInventory
// @Tags InInventory
// @Summary 批量删除InInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除InInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /In/deleteInInventoryByIds [delete]
func (InApi *InInventoryApi) DeleteInInventoryByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := InService.DeleteInInventoryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateInInventory 更新InInventory
// @Tags InInventory
// @Summary 更新InInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventory.InInventory true "更新InInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /In/updateInInventory [put]
func (InApi *InInventoryApi) UpdateInInventory(c *gin.Context) {
	var In inventory.InInventory
	err := c.ShouldBindJSON(&In)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Date":{utils.NotEmpty()},
          "GoodType":{utils.NotEmpty()},
          "Number":{utils.NotEmpty()},
          "Unit":{utils.NotEmpty()},
          "Operator":{utils.NotEmpty()},
      }
    if err := utils.Verify(In, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := InService.UpdateInInventory(In); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindInInventory 用id查询InInventory
// @Tags InInventory
// @Summary 用id查询InInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query inventory.InInventory true "用id查询InInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /In/findInInventory [get]
func (InApi *InInventoryApi) FindInInventory(c *gin.Context) {
	var In inventory.InInventory
	err := c.ShouldBindQuery(&In)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reIn, err := InService.GetInInventory(In.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reIn": reIn}, c)
	}
}

// GetInInventoryList 分页获取InInventory列表
// @Tags InInventory
// @Summary 分页获取InInventory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query inventoryReq.InInventorySearch true "分页获取InInventory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /In/getInInventoryList [get]
func (InApi *InInventoryApi) GetInInventoryList(c *gin.Context) {
	var pageInfo inventoryReq.InInventorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := InService.GetInInventoryInfoList(pageInfo); err != nil {
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
