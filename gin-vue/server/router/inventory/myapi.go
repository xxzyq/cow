package inventory

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type MyApi struct {
}

func (m *MyApi) InitMyApiRouter(Router *gin.RouterGroup) {
	MyRouterGroup := Router.Group("api")
	api := v1.ApiGroupApp.InventoryApiGroup.MyApi
	{
		MyRouterGroup.POST("create", api.CreateApi)
	}
}
