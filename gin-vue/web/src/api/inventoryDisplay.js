import service from '@/utils/request'

// @Tags InventoryDisplay
// @Summary 创建InventoryDisplay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InventoryDisplay true "创建InventoryDisplay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inventDis/createInventoryDisplay [post]
export const createInventoryDisplay = (data) => {
  return service({
    url: '/inventDis/createInventoryDisplay',
    method: 'post',
    data
  })
}

// @Tags InventoryDisplay
// @Summary 删除InventoryDisplay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InventoryDisplay true "删除InventoryDisplay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inventDis/deleteInventoryDisplay [delete]
export const deleteInventoryDisplay = (data) => {
  return service({
    url: '/inventDis/deleteInventoryDisplay',
    method: 'delete',
    data
  })
}

// @Tags InventoryDisplay
// @Summary 删除InventoryDisplay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除InventoryDisplay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inventDis/deleteInventoryDisplay [delete]
export const deleteInventoryDisplayByIds = (data) => {
  return service({
    url: '/inventDis/deleteInventoryDisplayByIds',
    method: 'delete',
    data
  })
}

// @Tags InventoryDisplay
// @Summary 删除InventoryDisplay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除InventoryDisplay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inventDis/deleteInventoryDisplay [delete]
export const deleteInventoryDisplayByGoodType = (data) => {
  return service({
    url: '/inventDis/deleteInventoryDisplayByGoodType',
    method: 'delete',
    data
  })
}

// @Tags InventoryDisplay
// @Summary 更新InventoryDisplay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InventoryDisplay true "更新InventoryDisplay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /inventDis/updateInventoryDisplay [put]
export const updateInventoryDisplay = (data) => {
  return service({
    url: '/inventDis/updateInventoryDisplay',
    method: 'put',
    data
  })
}

// @Tags InventoryDisplay
// @Summary 用id查询InventoryDisplay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.InventoryDisplay true "用id查询InventoryDisplay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /inventDis/findInventoryDisplay [get]
export const findInventoryDisplay = (params) => {
  return service({
    url: '/inventDis/findInventoryDisplay',
    method: 'get',
    params
  })
}

// @Tags InventoryDisplay
// @Summary 用id查询InventoryDisplay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.InventoryDisplay true "用id查询InventoryDisplay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /inventDis/findInventoryDisplay [get]
export const findInventoryDisplayByGoodType = (params) => {
  return service({
    url: '/inventDis/findInventoryDisplayByGoodType',
    method: 'get',
    params
  })
}

// @Tags InventoryDisplay
// @Summary 分页获取InventoryDisplay列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取InventoryDisplay列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inventDis/getInventoryDisplayList [get]
export const getInventoryDisplayList = (params) => {
  return service({
    url: '/inventDis/getInventoryDisplayList',
    method: 'get',
    params
  })
}
