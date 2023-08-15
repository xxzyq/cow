package inventory

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type MyApi struct {
}

func (m *MyApi) CreateApi(c *gin.Context) {
	myApiService.CreateApiS()
	response.Ok(c)
}
