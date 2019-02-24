package router

import (
	routerV1 "github.com/fajfar-com/confix.server/internal/router/v1"
	"github.com/gin-gonic/gin"
)

func GetRoutes(router *gin.Engine) *gin.Engine {

	routerV1.GetRouters(router)

	return router
}