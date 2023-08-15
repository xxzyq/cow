import service from '@/utils/request'

// @Tags OutInventory
// @Summary 创建OutInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OutInventory true "创建OutInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /out/createOutInventory [post]
export const createOutInventory = (data) => {
  return service({
    url: '/out/createOutInventory',
    method: 'post',
    data
  })
}

// @Tags OutInventory
// @Summary 删除OutInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OutInventory true "删除OutInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /out/deleteOutInventory [delete]
export const deleteOutInventory = (data) => {
  return service({
    url: '/out/deleteOutInventory',
    method: 'delete',
    data
  })
}

// @Tags OutInventory
// @Summary 删除OutInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OutInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /out/deleteOutInventory [delete]
export const deleteOutInventoryByIds = (data) => {
  return service({
    url: '/out/deleteOutInventoryByIds',
    method: 'delete',
    data
  })
}

// @Tags OutInventory
// @Summary 更新OutInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OutInventory true "更新OutInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /out/updateOutInventory [put]
export const updateOutInventory = (data) => {
  return service({
    url: '/out/updateOutInventory',
    method: 'put',
    data
  })
}

// @Tags OutInventory
// @Summary 用id查询OutInventory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OutInventory true "用id查询OutInventory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /out/findOutInventory [get]
export const findOutInventory = (params) => {
  return service({
    url: '/out/findOutInventory',
    method: 'get',
    params
  })
}

// @Tags OutInventory
// @Summary 分页获取OutInventory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OutInventory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /out/getOutInventoryList [get]
export const getOutInventoryList = (params) => {
  return service({
    url: '/out/getOutInventoryList',
    method: 'get',
    params
  })
}
