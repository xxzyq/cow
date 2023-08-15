package weightRecord

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WeightRecordRouter struct {
}

// InitWeightRecordRouter 初始化 WeightRecord 路由信息
func (s *WeightRecordRouter) InitWeightRecordRouter(Router *gin.RouterGroup) {
	weightRouter := Router.Group("weight").Use(middleware.OperationRecord())
	weightRouterWithoutRecord := Router.Group("weight")
	var weightApi = v1.ApiGroupApp.WeightRecordApiGroup.WeightRecordApi
	{
		weightRouter.POST("createWeightRecord", weightApi.CreateWeightRecord)   // 新建WeightRecord
		weightRouter.DELETE("deleteWeightRecord", weightApi.DeleteWeightRecord) // 删除WeightRecord
		weightRouter.DELETE("deleteWeightRecordByIds", weightApi.DeleteWeightRecordByIds) // 批量删除WeightRecord
		weightRouter.PUT("updateWeightRecord", weightApi.UpdateWeightRecord)    // 更新WeightRecord
	}
	{
		weightRouterWithoutRecord.GET("findWeightRecord", weightApi.FindWeightRecord)        // 根据ID获取WeightRecord
		weightRouterWithoutRecord.GET("getWeightRecordList", weightApi.GetWeightRecordList)  // 获取WeightRecord列表
	}
}
