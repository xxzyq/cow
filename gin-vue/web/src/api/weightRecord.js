import service from '@/utils/request'

// @Tags WeightRecord
// @Summary 创建WeightRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WeightRecord true "创建WeightRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weight/createWeightRecord [post]
export const createWeightRecord = (data) => {
  return service({
    url: '/weight/createWeightRecord',
    method: 'post',
    data
  })
}

// @Tags WeightRecord
// @Summary 删除WeightRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WeightRecord true "删除WeightRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weight/deleteWeightRecord [delete]
export const deleteWeightRecord = (data) => {
  return service({
    url: '/weight/deleteWeightRecord',
    method: 'delete',
    data
  })
}

// @Tags WeightRecord
// @Summary 删除WeightRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WeightRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weight/deleteWeightRecord [delete]
export const deleteWeightRecordByIds = (data) => {
  return service({
    url: '/weight/deleteWeightRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags WeightRecord
// @Summary 更新WeightRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WeightRecord true "更新WeightRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /weight/updateWeightRecord [put]
export const updateWeightRecord = (data) => {
  return service({
    url: '/weight/updateWeightRecord',
    method: 'put',
    data
  })
}

// @Tags WeightRecord
// @Summary 用id查询WeightRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WeightRecord true "用id查询WeightRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /weight/findWeightRecord [get]
export const findWeightRecord = (params) => {
  return service({
    url: '/weight/findWeightRecord',
    method: 'get',
    params
  })
}

// @Tags WeightRecord
// @Summary 分页获取WeightRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取WeightRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weight/getWeightRecordList [get]
export const getWeightRecordList = (params) => {
  return service({
    url: '/weight/getWeightRecordList',
    method: 'get',
    params
  })
}
