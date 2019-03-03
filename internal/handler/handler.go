package handler

import (
	"github.com/fajfar-com/confix.server/internal/configMng"
	"github.com/gin-gonic/gin"
)

func init(){

}

func GetPing() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func GetApplicationConfiguration() gin.HandlerFunc {
	return func(context *gin.Context) {
		var appName = context.Param("app")
		var appConfig = configMng.GetApplication(appName)

		context.JSON(200, appConfig.Configs)
	}
}