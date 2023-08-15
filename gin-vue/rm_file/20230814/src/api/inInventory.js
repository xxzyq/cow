import service from '@/utils/request'

// @Tags InInventory
// @Summary 创建InInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InInventory true "创建InInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /In/createInInventory [post]
export const createInInventory = (data) => {

  return service({
    url: '/In/createInInventory',
    method: 'post',
    data
  })
}

// @Tags InInventory
// @Summary 删除InInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InInventory true "删除InInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /In/deleteInInventory [delete]
export const deleteInInventory = (data) => {
  return service({
    url: '/In/deleteInInventory',
    method: 'delete',
    data
  })
}

// @Tags InInventory
// @Summary 删除InInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除InInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /In/deleteInInventory [delete]
export const deleteInInventoryByIds = (data) => {
  return service({
    url: '/In/deleteInInventoryByIds',
    method: 'delete',
    data
  })
}

// @Tags InInventory
// @Summary 更新InInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InInventory true "更新InInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /In/updateInInventory [put]
export const updateInInventory = (data) => {
  return service({
    url: '/In/updateInInventory',
    method: 'put',
    data
  })
}

// @Tags InInventory
// @Summary 用id查询InInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.InInventory true "用id查询InInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /In/findInInventory [get]
export const findInInventory = (params) => {
  return service({
    url: '/In/findInInventory',
    method: 'get',
    params
  })
}

// @Tags InInventory
// @Summary 分页获取InInventory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取InInventory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /In/getInInventoryList [get]
export const getInInventoryList = (params) => {
  return service({
    url: '/In/getInInventoryList',
    method: 'get',
    params
  })
}
