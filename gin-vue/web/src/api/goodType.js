import service from '@/utils/request'
// @Tags SysApi
// @Summary 创建商品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "创建商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [post]
export const createExaGoodType = (data) => {
  return service({
    url: '/goodType/goodType',
    method: 'post',
    data
  })
}



// @Tags SysApi
// @Summary 删除库存类型信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "删除库存类型信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [delete]
export const deleteExaGoodType = (data) => {
  return service({
    url: '/goodType/goodType',
    method: 'delete',
    data
  })
}

// @Tags SysApi
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "获取单一客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [get]
export const getExaGoodType = (params) => {
    return service({
      url: '/goodType/goodType',
      method: 'get',
      params
    })
  }
  
  // @Tags SysApi
  // @Summary 获取权限客户列表
  // @Security ApiKeyAuth
  // @accept application/json
  // @Produce application/json
  // @Param data body modelInterface.PageInfo true "获取权限客户列表"
  // @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
  // @Router /customer/customerList [get]
  export const getExaGoodTypeList = (params) => {
    return service({
      url: '/goodType/goodType',
      method: 'get',
      params
    })
  }
