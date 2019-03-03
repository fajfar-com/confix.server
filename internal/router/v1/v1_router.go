package v1

import (
	"github.com/fajfar-com/confix.server/internal/handler"
	"github.com/gin-gonic/gin"
)

func GetRouters(router *gin.Engine) *gin.Engine  {

	group := router.Group("/v1")
	{
		group.GET("/ping", handler.GetPing())
		group.GET("/applications/:app", handler.GetApplicationConfiguration())
	}

	return router
}