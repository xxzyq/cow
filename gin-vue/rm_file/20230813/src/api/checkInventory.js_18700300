import service from '@/utils/request'

// @Tags CheckInventory
// @Summary 创建CheckInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckInventory true "创建CheckInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /check/createCheckInventory [post]
export const createCheckInventory = (data) => {
  return service({
    url: '/check/createCheckInventory',
    method: 'post',
    data
  })
}

// @Tags CheckInventory
// @Summary 删除CheckInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckInventory true "删除CheckInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /check/deleteCheckInventory [delete]
export const deleteCheckInventory = (data) => {
  return service({
    url: '/check/deleteCheckInventory',
    method: 'delete',
    data
  })
}

// @Tags CheckInventory
// @Summary 删除CheckInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CheckInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /check/deleteCheckInventory [delete]
export const deleteCheckInventoryByIds = (data) => {
  return service({
    url: '/check/deleteCheckInventoryByIds',
    method: 'delete',
    data
  })
}

// @Tags CheckInventory
// @Summary 更新CheckInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CheckInventory true "更新CheckInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /check/updateCheckInventory [put]
export const updateCheckInventory = (data) => {
  return service({
    url: '/check/updateCheckInventory',
    method: 'put',
    data
  })
}

// @Tags CheckInventory
// @Summary 用id查询CheckInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CheckInventory true "用id查询CheckInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /check/findCheckInventory [get]
export const findCheckInventory = (params) => {
  return service({
    url: '/check/findCheckInventory',
    method: 'get',
    params
  })
}

// @Tags CheckInventory
// @Summary 分页获取CheckInventory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CheckInventory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /check/getCheckInventoryList [get]
export const getCheckInventoryList = (params) => {
  return service({
    url: '/check/getCheckInventoryList',
    method: 'get',
    params
  })
}
