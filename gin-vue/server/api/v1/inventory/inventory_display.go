package inventory

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/inventory"
	inventoryReq "github.com/flipped-aurora/gin-vue-admin/server/model/inventory/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

type InventoryDisplayApi struct {
}

var inventDisService = service.ServiceGroupApp.InventoryServiceGroup.InventoryDisplayService

// CreateInventoryDisplay 创建InventoryDisplay
// @Tags InventoryDisplay
// @Summary 创建InventoryDisplay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventory.InventoryDisplay true "创建InventoryDisplay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inventDis/createInventoryDisplay [post]
func (inventDisApi *InventoryDisplayApi) CreateInventoryDisplay(c *gin.Context) {
	var inventDis inventory.InventoryDisplay
	err := c.ShouldBindJSON(&inventDis)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"GoodType": {utils.NotEmpty()},
		"Number":   {utils.NotEmpty()},
		"Unit":     {utils.NotEmpty()},
	}
	if err := utils.Verify(inventDis, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := inventDisService.CreateInventoryDisplay(&inventDis); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteInventoryDisplay 删除InventoryDisplay
// @Tags InventoryDisplay
// @Summary 删除InventoryDisplay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventory.InventoryDisplay true "删除InventoryDisplay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inventDis/deleteInventoryDisplay [delete]
func (inventDisApi *InventoryDisplayApi) DeleteInventoryDisplay(c *gin.Context) {
	var inventDis inventory.InventoryDisplay
	err := c.ShouldBindJSON(&inventDis)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := inventDisService.DeleteInventoryDisplay(inventDis); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteInventoryDisplayByIds 批量删除InventoryDisplay
// @Tags InventoryDisplay
// @Summary 批量删除InventoryDisplay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除InventoryDisplay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /inventDis/deleteInventoryDisplayByIds [delete]
func (inventDisApi *InventoryDisplayApi) DeleteInventoryDisplayByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := inventDisService.DeleteInventoryDisplayByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// DeleteInventoryDisplayByIds 批量删除InventoryDisplay
// @Tags InventoryDisplay
// @Summary 批量删除InventoryDisplay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除InventoryDisplay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /inventDis/deleteInventoryDisplayByIds [delete]
func (inventDisApi *InventoryDisplayApi) DeleteInventoryDisplayByGoodType(c *gin.Context) {
	var GT request.GoodTypeReq
	err := c.ShouldBindJSON(&GT)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := inventDisService.DeleteInventoryDisplayByGoodType(GT); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateInventoryDisplay 更新InventoryDisplay
// @Tags InventoryDisplay
// @Summary 更新InventoryDisplay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventory.InventoryDisplay true "更新InventoryDisplay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /inventDis/updateInventoryDisplay [put]
func (inventDisApi *InventoryDisplayApi) UpdateInventoryDisplay(c *gin.Context) {
	var inventDis inventory.InventoryDisplay
	err := c.ShouldBindJSON(&inventDis)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"GoodType": {utils.NotEmpty()},
		"Number":   {utils.NotEmpty()},
		"Unit":     {utils.NotEmpty()},
	}
	if err := utils.Verify(inventDis, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := inventDisService.UpdateInventoryDisplay(inventDis); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindInventoryDisplay 用id查询InventoryDisplay
// @Tags InventoryDisplay
// @Summary 用id查询InventoryDisplay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query inventory.InventoryDisplay true "用id查询InventoryDisplay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /inventDis/findInventoryDisplay [get]
func (inventDisApi *InventoryDisplayApi) FindInventoryDisplay(c *gin.Context) {
	var inventDis inventory.InventoryDisplay
	err := c.ShouldBindQuery(&inventDis)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reinventDis, err := inventDisService.GetInventoryDisplay(inventDis.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinventDis": reinventDis}, c)
	}
}

func (inventDisApi *InventoryDisplayApi) FindInventoryDisplayByGoodType(c *gin.Context) {
	var inventDis inventory.InventoryDisplay
	err := c.ShouldBindQuery(&inventDis)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reinventDis, err := inventDisService.GetInventoryDisplayByGoodType(inventDis.GoodType); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinventDis": reinventDis}, c)
	}
}

// GetInventoryDisplayList 分页获取InventoryDisplay列表
// @Tags InventoryDisplay
// @Summary 分页获取InventoryDisplay列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query inventoryReq.InventoryDisplaySearch true "分页获取InventoryDisplay列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inventDis/getInventoryDisplayList [get]
func (inventDisApi *InventoryDisplayApi) GetInventoryDisplayList(c *gin.Context) {
	var pageInfo inventoryReq.InventoryDisplaySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := inventDisService.GetInventoryDisplayInfoList(pageInfo); err != nil {
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

type InventoryDisplay struct{}

// DeleteInventoryDisplay receives a JSON object and deletes the corresponding record from the database.
func DeleteInventoryDisplay(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var requestData map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	goodType, ok := requestData["goodType"].(string)
	if !ok {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	var display InventoryDisplay
	if err := db.Where("good_type = ?", goodType).First(&display).Error; err != nil {
		http.Error(w, fmt.Sprintf("Error finding record: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if err := db.Delete(&display).Error; err != nil {
		http.Error(w, fmt.Sprintf("Error deleting record: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Record deleted successfully"))
}
